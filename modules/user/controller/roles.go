package controller

import (
	"fmt"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/user/dto"
	"github.com/atom-apps/door/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type RoleController struct {
	roleSvc *service.RoleService
}

// LabelShow
//
//	@Summary		LabelShow
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{object}	dto.UserItem
//	@Router			/v1/users/roles/{id}/label [get]
func (c *RoleController) LabelShow(ctx *fiber.Ctx, id int64) ([]common.LabelItem, error) {
	item, err := c.roleSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return []common.LabelItem{
		{Label: "ID", Value: fmt.Sprintf("%d", item.ID)},
		{Label: "名称", Value: item.Name},
		{Label: "别名", Value: item.Slug},
		{Label: "父ID", Value: fmt.Sprintf("%d", item.ParentID)},
	}, nil
}

// Filters get list filter items
//
//	@Summary		get list filters
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200			{array}	common.Filter
//	@Router			/v1/users/roles/filters [get]
func (c *RoleController) Filters(ctx *fiber.Ctx) ([]common.Filter, error) {
	return dto.RoleListQueryFilters(), nil
}

// Columns of list
//
//	@Summary		get list columns
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	common.Columns
//	@Router			/v1/users/roles/columns [get]
func (c *RoleController) Columns(ctx *fiber.Ctx) (common.Columns, error) {
	columns := []common.TableColumnData{
		{Title: "ID", DataIndex: "id", Hidden: true},
		{Title: "名称", DataIndex: "name"},
		{Title: "别名", DataIndex: "slug"},
		{Title: "描述", DataIndex: "description"},
		{Title: "父级", DataIndex: "parent_id"},
		{Title: "操作", DataIndex: "operations", Align: lo.ToPtr("right")},
	}

	return common.Columns{
		Columns: lo.Map(columns, func(item common.TableColumnData, _ int) common.TableColumnData {
			return item.Format()
		}),
		Hidden: lo.FilterMap(columns, func(item common.TableColumnData, _ int) (string, bool) {
			if item.Hidden {
				return item.DataIndex, true
			}
			return "", false
		}),
	}, nil
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"RoleID"
//	@Success		200	{object}	dto.RoleItem
//	@Router			/v1/users/roles/{id} [get]
func (c *RoleController) Show(ctx *fiber.Ctx, id int64) (*dto.RoleItem, error) {
	item, err := c.roleSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.roleSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.RoleListQueryFilter	true	"RoleListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.RoleItem}
//	@Router			/v1/users/roles [get]
func (c *RoleController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.RoleListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.roleSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.roleSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.RoleForm	true	"RoleForm"
//	@Success		200		{string}	RoleID
//	@Router			/v1/users/roles [post]
func (c *RoleController) Create(ctx *fiber.Ctx, body *dto.RoleForm) error {
	return c.roleSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"RoleID"
//	@Param			body	body		dto.RoleForm	true	"RoleForm"
//	@Success		200		{string}	RoleID
//	@Failure		500		{string}	RoleID
//	@Router			/v1/users/roles/{id} [put]
func (c *RoleController) Update(ctx *fiber.Ctx, id int64, body *dto.RoleForm) error {
	return c.roleSvc.Update(ctx.Context(), id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"RoleID"
//	@Success		200	{string}	RoleID
//	@Failure		500	{string}	RoleID
//	@Router			/v1/users/roles/{id} [delete]
func (c *RoleController) Delete(ctx *fiber.Ctx, id int64) error {
	return c.roleSvc.Delete(ctx.Context(), id)
}
