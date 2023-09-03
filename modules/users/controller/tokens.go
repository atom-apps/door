package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-apps/door/modules/users/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type TokenController struct {
	tokenSvc *service.TokenService
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
//	@Router			/v1/users/tokens [get]
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
//	@Router			/v1/users/tokens/{id} [delete]
func (c *TokenController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.tokenSvc.Delete(ctx.Context(), id)
}
