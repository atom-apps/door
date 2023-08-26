package service

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/modules/systems/dao"
	"github.com/atom-apps/door/modules/systems/dto"
	"github.com/samber/lo"

	"github.com/jinzhu/copier"
)

// @provider
type RouteService struct {
	routeDao *dao.RouteDao
}

func (svc *RouteService) DecorateItem(model *models.Route, id int) *dto.RouteItem {
	dtoItem := &dto.RouteItem{
		ID:       model.ID,
		Type:     &model.Type,
		ParentID: model.ParentID,
		Name:     model.Name,
		Path:     model.Path,
		Metadata: model.Metadata.Data,
		Children: []*dto.RouteItem{},
	}

	return dtoItem
}

func (svc *RouteService) Tree(ctx context.Context, mode consts.RouteType, parentID int64) ([]*dto.RouteItem, error) {
	items, err := svc.routeDao.FindByParentIDOfMode(ctx, mode, parentID)
	if err != nil {
		return nil, err
	}

	return lo.Map(items, func(model *models.Route, index int) *dto.RouteItem {
		dtoItem := &dto.RouteItem{
			ID:       model.ID,
			Type:     &model.Type,
			ParentID: model.ParentID,
			Name:     model.Name,
			Path:     model.Path,
			Metadata: model.Metadata.Data,
			Children: []*dto.RouteItem{},
		}

		item, err := svc.Tree(ctx, mode, model.ID)
		if err == nil {
			dtoItem.Children = item
		}

		return dtoItem
	}), nil
}

func (svc *RouteService) GetByID(ctx context.Context, id int64) (*models.Route, error) {
	return svc.routeDao.GetByID(ctx, id)
}

func (svc *RouteService) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RouteListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Route, error) {
	return svc.routeDao.FindByQueryFilter(ctx, queryFilter, sortFilter)
}

func (svc *RouteService) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.RouteListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Route, int64, error) {
	return svc.routeDao.PageByQueryFilter(ctx, queryFilter, pageFilter.Format(), sortFilter)
}

// CreateFromModel
func (svc *RouteService) CreateFromModel(ctx context.Context, model *models.Route) error {
	return svc.routeDao.Create(ctx, model)
}

// Create
func (svc *RouteService) Create(ctx context.Context, body *dto.RouteForm) error {
	model := &models.Route{}
	_ = copier.Copy(model, body)
	return svc.routeDao.Create(ctx, model)
}

// Update
func (svc *RouteService) Update(ctx context.Context, id int64, body *dto.RouteForm) error {
	model, err := svc.GetByID(ctx, id)
	if err != nil {
		return err
	}

	_ = copier.Copy(model, body)
	model.ID = id
	return svc.routeDao.Update(ctx, model)
}

// UpdateFromModel
func (svc *RouteService) UpdateFromModel(ctx context.Context, model *models.Route) error {
	return svc.routeDao.Update(ctx, model)
}

// Delete
func (svc *RouteService) Delete(ctx context.Context, id int64) error {
	return svc.routeDao.Delete(ctx, id)
}
