package controller

import (
	"github.com/atom-apps/door/modules/users/dto"
	"github.com/atom-apps/door/modules/users/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type UserAddressController struct {
	userAddressSvc *service.UserAddressService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			用户地址
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int	true	"UserId"
//	@Param			id	path		int	true	"UserAddressID"
//	@Success		200	{object}	dto.UserAddressItem
//	@Router			/v1/users/users/{user_id}/addresses/{id} [get]
func (c *UserAddressController) Show(ctx *fiber.Ctx, userId, id uint64) (*dto.UserAddressItem, error) {
	item, err := c.userAddressSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.userAddressSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			用户地址
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int	true	"UserId"
//	@Success		200			{array}	dto.UserAddressItem
//	@Router			/v1/users/users/{user_id}/addresses [get]
func (c *UserAddressController) List(ctx *fiber.Ctx, userID uint64) ([]*dto.UserAddressItem, error) {
	items, err := c.userAddressSvc.FindByUserID(ctx.Context(), userID)
	if err != nil {
		return nil, err
	}
	return lo.Map(items, c.userAddressSvc.DecorateItem), nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			用户地址
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int	true	"UserId"
//	@Param			body	body		dto.UserAddressForm	true	"UserAddressForm"
//	@Success		200		{string}	UserAddressID
//	@Router			/v1/users/users/{user_id}/addresses [post]
func (c *UserAddressController) Create(ctx *fiber.Ctx, userId uint64, body *dto.UserAddressForm) error {
	return c.userAddressSvc.Create(ctx.Context(), userId, body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			用户地址
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int	true	"UserId"
//	@Param			id		path		int				true	"UserAddressID"
//	@Param			body	body		dto.UserAddressForm	true	"UserAddressForm"
//	@Success		200		{string}	UserAddressID
//	@Router			/v1/users/users/{user_id}/addresses/{id} [put]
func (c *UserAddressController) Update(ctx *fiber.Ctx, userId, id uint64, body *dto.UserAddressForm) error {
	return c.userAddressSvc.Update(ctx.Context(), userId, id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			用户地址
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		int	true	"UserId"
//	@Param			id	path		int	true	"UserAddressID"
//	@Success		200	{string}	UserAddressID
//	@Router			/v1/users/users/{user_id}/addresses/{id} [delete]
func (c *UserAddressController) Delete(ctx *fiber.Ctx, userId, id uint64) error {
	return c.userAddressSvc.Delete(ctx.Context(), userId, id)
}
