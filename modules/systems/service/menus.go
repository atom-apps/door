package service

import (
	"context"
	"fmt"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/systems/dao"
	"github.com/atom-apps/door/modules/systems/dto"
	"github.com/samber/lo"
)

// @provider
type MenuService struct {
	menuDao *dao.MenuDao
}

func (svc *MenuService) DecorateItem(model *models.Menu, id int) *dto.MenuItem {
	return &dto.MenuItem{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		Name:      model.Name,
		Slug:      model.Slug,
		GroupID:   model.GroupID,
		ParentID:  model.ParentID,
		Metadata:  model.Metadata,
	}
}

func (svc *MenuService) GetByID(ctx context.Context, id uint64) (*models.Menu, error) {
	return svc.menuDao.GetByID(ctx, id)
}

func (svc *MenuService) GetGroupByID(ctx context.Context, groupID uint64) (*models.Menu, error) {
	return svc.menuDao.GetGroupByID(ctx, groupID)
}

func (svc *MenuService) GetGroupTree(ctx context.Context, groupID uint64) ([]*dto.MenuTreeItem, error) {
	items, err := svc.menuDao.GetGroupItemsByID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	return svc.genTree(items, groupID), nil
}

func (svc *MenuService) genTree(items []*models.Menu, parentID uint64) []*dto.MenuTreeItem {
	var tree []*dto.MenuTreeItem
	for _, item := range items {
		if item.ParentID == parentID {
			tree = append(tree, &dto.MenuTreeItem{
				Key:      fmt.Sprintf("%d", item.ID),
				Title:    item.Name,
				Children: svc.genTree(items, item.ID),
			})
		}
	}
	return tree
}

func (svc *MenuService) PageGroupByQueryFilter(ctx context.Context, queryFilter *dto.MenuListQueryFilter, pageFilter *common.PageQueryFilter, sortFilter *common.SortQueryFilter) ([]*models.Menu, int64, error) {
	queryFilter.GroupID = lo.ToPtr[uint64](0)
	return svc.menuDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

func (svc *MenuService) FindGroupByQueryFilter(ctx context.Context, queryFilter *dto.MenuListQueryFilter, sortFilter *common.SortQueryFilter) ([]*models.Menu, error) {
	queryFilter.GroupID = lo.ToPtr[uint64](0)
	return svc.menuDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

// CreateFromModel
func (svc *MenuService) CreateFromModel(ctx context.Context, model *models.Menu) error {
	return svc.menuDao.Create(ctx, model)
}

// Create
func (svc *MenuService) Create(ctx context.Context, body *dto.MenuForm, groupID uint64) error {
	model := &models.Menu{
		Name:     body.Name,
		Slug:     body.Slug,
		GroupID:  groupID,
		ParentID: body.ParentID,
		Metadata: body.Metadata,
	}
	return svc.menuDao.Create(ctx, model)
}

// Update
func (svc *MenuService) Update(ctx context.Context, id uint64, body *dto.MenuForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	model.Name = body.Name
	model.Slug = body.Slug
	model.ParentID = body.ParentID
	model.Metadata = body.Metadata

	return svc.menuDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *MenuService) UpdateFromModel(ctx context.Context, model *models.Menu) error {
	return svc.menuDao.Update(ctx, model)
}

// Delete
func (svc *MenuService) Delete(ctx context.Context, id uint64) error {
	item, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return svc.menuDao.Transaction(func() error {
		if err := svc.menuDao.Delete(ctx, id); err != nil {
			return err
		}
		items, err := svc.menuDao.GetGroupSubItemsByID(ctx, item.GroupID, id)
		if err != nil {
			return err
		}

		return svc.menuDao.Delete(ctx, lo.Map(items, func(item *models.Menu, _ int) uint64 {
			return item.ID
		})...)
	})
}
