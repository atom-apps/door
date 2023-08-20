package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/user/dto"
	"github.com/atom-apps/door/modules/user/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type TokenController struct {
	tokenSvc *service.TokenService
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"TokenID"
//	@Success		200	{object}	dto.TokenItem
//	@Router			/users/tokens/{id} [get]
func (c *TokenController) Show(ctx *fiber.Ctx, id int64) (*dto.TokenItem, error) {
	item, err := c.tokenSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.tokenSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.TokenListQueryFilter	true	"TokenListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter		true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter		true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.TokenItem}
//	@Router			/users/tokens [get]
func (c *TokenController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.TokenListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.tokenSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.tokenSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.TokenForm	true	"TokenForm"
//	@Success		200		{string}	TokenID
//	@Router			/users/tokens [post]
func (c *TokenController) Create(ctx *fiber.Ctx, body *dto.TokenForm) error {
	return c.tokenSvc.Create(ctx.Context(), body)
}

// Update update by id
//
//	@Summary		update by id
//	@Description	update by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"TokenID"
//	@Param			body	body		dto.TokenForm	true	"TokenForm"
//	@Success		200		{string}	TokenID
//	@Failure		500		{string}	TokenID
//	@Router			/users/tokens/{id} [put]
func (c *TokenController) Update(ctx *fiber.Ctx, id int64, body *dto.TokenForm) error {
	return c.tokenSvc.Update(ctx.Context(), id, body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"TokenID"
//	@Success		200	{string}	TokenID
//	@Failure		500	{string}	TokenID
//	@Router			/users/tokens/{id} [delete]
func (c *TokenController) Delete(ctx *fiber.Ctx, id int64) error {
	return c.tokenSvc.Delete(ctx.Context(), id)
}
