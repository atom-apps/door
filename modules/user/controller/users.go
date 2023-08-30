package controller

import (
	"fmt"
	"time"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/user/dto"
	"github.com/atom-apps/door/modules/user/service"
	"github.com/atom-providers/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type UserController struct {
	userSvc           *service.UserService
	permissionRuleSvc *service.PermissionRuleService
}

// Profile get current user info
//
//	@Summary		get current user info
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{object}	dto.UserItem
//	@Router			/v1/users/profile [get]
func (c *UserController) Profile(ctx *fiber.Ctx, claim *jwt.Claims) (*dto.UserItem, error) {
	item, err := c.userSvc.GetByID(ctx.Context(), claim.UserID)
	if err != nil {
		return nil, err
	}

	return c.userSvc.DecorateItem(item, 0), nil
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{object}	dto.UserItem
//	@Router			/v1/users/{id} [get]
func (c *UserController) Show(ctx *fiber.Ctx, id int64) (*dto.UserItem, error) {
	item, err := c.userSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Second * 2)

	return c.userSvc.DecorateItem(item, 0), nil
}

// LabelShow
//
//	@Summary		LabelShow
//	@Description	get info by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{object}	dto.UserItem
//	@Router			/v1/users/{id}/label [get]
func (c *UserController) LabelShow(ctx *fiber.Ctx, id int64) ([]common.LabelItem, error) {
	item, err := c.userSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return []common.LabelItem{
		{Label: "ID", Value: fmt.Sprintf("%d", item.ID)},
		{Label: "UUID", Value: item.UUID},
		{Label: "用户名", Value: item.Username},
		{Label: "昵称", Value: item.DisplayName},
		{Label: "电子邮箱", Value: item.Email},
		{Label: "邮箱验证", Value: common.BoolOrStr(item.EmailVerified, "已验证", "未验证")},
		{Label: "手机", Value: item.Phone},
		{Label: "状态", Value: item.Status.Cn()},
	}, nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.UserListQueryFilter	true	"UserListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.UserItem}
//	@Router			/v1/users [get]
func (c *UserController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.userSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.userSvc.DecorateItem),
	}, nil
}

// Filters get list filter items
//
//	@Summary		get list filters
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200			{array}	common.Filter
//	@Router			/v1/users/filters [get]
func (c *UserController) Filters(ctx *fiber.Ctx) ([]common.Filter, error) {
	return dto.UserListQueryFilters(), nil
}

// Columns of list
//
//	@Summary		get list columns
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200			{object}	common.Columns
//	@Router			/v1/users/columns [get]
func (c *UserController) Columns(ctx *fiber.Ctx) (common.Columns, error) {
	columns := []common.TableColumnData{
		{Title: "ID", DataIndex: "id"},
		{Title: "UUID", DataIndex: "uuid", Hidden: true},
		{Title: "昵称", DataIndex: "display_name"},
		{Title: "名称", DataIndex: "username"},
		{Title: "状态", Align: lo.ToPtr("center"), DataIndex: "status"},
		{Title: "Email", DataIndex: "email"},
		{Title: "电话", DataIndex: "phone"},
		{Title: "创建时间", DataIndex: "created_at", Hidden: true},
		{Title: "更新时间", DataIndex: "updated_at"},
		{Title: "操作", DataIndex: "operations", Align: lo.ToPtr("right")},
	}
	return common.NewColumns(columns), nil
}

// // Role users
// //
// //	@Summary		list by query filter
// //	@Description	list by query filter
// //	@Tags			User
// //	@Accept			json
// //	@Produce		json
// //	@Param			queryFilter	query		dto.UserListQueryFilter	true	"UserListQueryFilter"
// //	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
// //	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
// //	@Success		200			{object}	common.PageDataResponse{list=dto.UserItem}
// //	@Router			/v1/users/roles/{id} [get]
// func (c *UserController) Role(
// 	ctx *fiber.Ctx,
// 	id int64,
// 	queryFilter *dto.UserListQueryFilter,
// 	pageFilter *common.PageQueryFilter,
// 	sortFilter *common.SortQueryFilter,
// ) (*common.PageDataResponse, error) {
// 	var err error
// 	queryFilter.IDs, err = c.permissionRuleSvc.GetUserIDsOfRole(ctx.Context(), id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	items, total, err := c.userSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &common.PageDataResponse{
// 		PageQueryFilter: *pageFilter,
// 		Total:           total,
// 		Items:           lo.Map(items, c.userSvc.DecorateItem),
// 	}, nil
// }

// // Tenant users
// //
// //	@Summary		list by query filter
// //	@Description	list by query filter
// //	@Tags			User
// //	@Accept			json
// //	@Produce		json
// //	@Param			queryFilter	query		dto.UserListQueryFilter	true	"UserListQueryFilter"
// //	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
// //	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
// //	@Success		200			{object}	common.PageDataResponse{list=dto.UserItem}
// //	@Router			/v1/users/tenants/{id} [get]
// func (c *UserController) Tenant(
// 	ctx *fiber.Ctx,
// 	id int64,
// 	queryFilter *dto.UserListQueryFilter,
// 	pageFilter *common.PageQueryFilter,
// 	sortFilter *common.SortQueryFilter,
// ) (*common.PageDataResponse, error) {
// 	var err error
// 	queryFilter.IDs, err = c.permissionRuleSvc.GetUserIDsOfTenant(ctx.Context(), id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	items, total, err := c.userSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &common.PageDataResponse{
// 		PageQueryFilter: *pageFilter,
// 		Total:           total,
// 		Items:           lo.Map(items, c.userSvc.DecorateItem),
// 	}, nil
// }

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.UserForm	true	"UserForm"
//	@Success		200		{string}	UserID
//	@Router			/v1/users [post]
func (c *UserController) Create(ctx *fiber.Ctx, body *dto.UserForm) error {
	return c.userSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"UserID"
//	@Param			body	body		dto.UserForm	true	"UserForm"
//	@Success		200		{string}	UserID
//	@Failure		500		{string}	UserID
//	@Router			/v1/users/{id} [put]
func (c *UserController) Update(ctx *fiber.Ctx, id int64, body *dto.UserForm) error {
	return c.userSvc.Update(ctx.Context(), id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserID"
//	@Success		200	{string}	UserID
//	@Failure		500	{string}	UserID
//	@Router			/v1/users/{id} [delete]
func (c *UserController) Delete(ctx *fiber.Ctx, id int64) error {
	return c.userSvc.Delete(ctx.Context(), id)
}
