package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-apps/door/modules/users/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type SessionController struct {
	sessionSvc *service.SessionService
}

// Show get single item info
//
//	@Summary		get by id
//	@Description	get info by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"SessionID"
//	@Success		200	{object}	dto.SessionItem
//	@Router			/v1/users/sessions/{id} [get]
func (c *SessionController) Show(ctx *fiber.Ctx, id uint64) (*dto.SessionItem, error) {
	item, err := c.sessionSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.sessionSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		list by query filter
//	@Description	list by query filter
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			queryFilter	query		dto.SessionListQueryFilter	true	"SessionListQueryFilter"
//	@Param			pageFilter	query		common.PageQueryFilter		true	"PageQueryFilter"
//	@Param			sortFilter	query		common.SortQueryFilter		true	"SortQueryFilter"
//	@Success		200			{object}	common.PageDataResponse{list=dto.SessionItem}
//	@Router			/v1/users/sessions [get]
func (c *SessionController) List(
	ctx *fiber.Ctx,
	queryFilter *dto.SessionListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*common.PageDataResponse, error) {
	items, total, err := c.sessionSvc.PageByQueryFilter(ctx.Context(), queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &common.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.sessionSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		create new item
//	@Description	create new item
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dto.SessionForm	true	"SessionForm"
//	@Success		200		{string}	SessionID
//	@Router			/v1/users/sessions [post]
func (c *SessionController) Create(ctx *fiber.Ctx, body *dto.SessionForm) error {
	return c.sessionSvc.Create(ctx.Context(), body)
}

// Delete delete by id
//
//	@Summary		delete by id
//	@Description	delete by id
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"SessionID"
//	@Router			/v1/users/sessions/{id} [delete]
func (c *SessionController) Delete(ctx *fiber.Ctx, id uint64) error {
	return c.sessionSvc.DeleteBySessionID(ctx.Context(), id)
}

// Delete delete by session string uuid
//
//	@Summary		delete by session uuid
//	@Description	delete by session uuid
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			sess_id	path	string	true	"SessionID"
//	@Router			/v1/users/sessions/{sess_id}/by-session-id [delete]
func (c *SessionController) DeleteBySessionID(ctx *fiber.Ctx, sessID string) error {
	m, err := c.sessionSvc.GetBySessionID(ctx.Context(), sessID)
	if err != nil {
		return err
	}
	return c.sessionSvc.DeleteBySessionID(ctx.Context(), m.ID)
}
