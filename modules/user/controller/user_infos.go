package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/user/dto"
	"github.com/atom-apps/door/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type UserInfoController struct {
	userInfoSvc *service.UserInfoService
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserInfoID"
//	@Success		200	{object}	dto.UserInfoItem
//	@Router			/user_infos/{id} [get]
func (c *UserInfoController) Show(ctx *fiber.Ctx, id int64) (*dto.UserInfoItem, error) {
	item, err := c.userInfoSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.userInfoSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.UserInfoListQueryFilter	true	"UserInfoListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.UserInfoItem}
//	@Router			/user_infos [get]
func (c *UserInfoController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.UserInfoListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.userInfoSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.userInfoSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.UserInfoForm	true	"UserInfoForm"
//	@Success		200		{string}	UserInfoID
//	@Router			/user_infos [post]
func (c *UserInfoController) Create(ctx *fiber.Ctx, body *dto.UserInfoForm) error {
	return c.userInfoSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"UserInfoID"
//	@Param			body	body		dto.UserInfoForm	true	"UserInfoForm"
//	@Success		200		{string}	UserInfoID
//	@Failure		500		{string}	UserInfoID
//	@Router			/user_infos/{id} [put]
func (c *UserInfoController) Update(ctx *fiber.Ctx, id int64, body *dto.UserInfoForm) error {
	return c.userInfoSvc.Update(ctx.Context(), id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"UserInfoID"
//	@Success		200	{string}	UserInfoID
//	@Failure		500	{string}	UserInfoID
//	@Router			/user_infos/{id} [delete]
func (c *UserInfoController) Delete(ctx *fiber.Ctx, id int64) error {
	return c.userInfoSvc.Delete(ctx.Context(), id)
}
