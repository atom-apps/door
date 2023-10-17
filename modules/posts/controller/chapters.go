package controller

import (
	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/modules/posts/dto"
	"github.com/atom-apps/door/modules/posts/service"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

// @provider
type ChapterController struct {
	chapterSvc *service.ChapterService
}

// Show get single item info
//
//	@Summary		Show
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			bookId	path		int	true	"BookId"
//	@Param			id	path		int	true	"ChapterID"
//	@Success		200	{object}	dto.ChapterItem
//	@Router			/v1/posts/books/{book_id}/chapters/{id} [get]
func (c *ChapterController) Show(ctx *fiber.Ctx, bookId, id uint64) (*dto.ChapterItem, error) {
	item, err := c.chapterSvc.GetByID(ctx.Context(), id)
	if err != nil {
		return nil, err
	}

	return c.chapterSvc.DecorateItem(item, 0), nil
}

// List list by query filter
//
//	@Summary		List
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			bookId	path		int	true	"BookId"
//	@Param			queryFilter	query		dto.ChapterListQueryFilter	true	"ChapterListQueryFilter"
//	@Param			pageFilter	query		ds.PageQueryFilter	true	"PageQueryFilter"
//	@Param			sortFilter	query		ds.SortQueryFilter	true	"SortQueryFilter"
//	@Success		200			{object}	ds.PageDataResponse{list=dto.ChapterItem}
//	@Router			/v1/posts/books/{book_id}/chapters [get]
func (c *ChapterController) List(
	ctx *fiber.Ctx,
	bookId uint64,
	queryFilter *dto.ChapterListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*ds.PageDataResponse, error) {
	items, total, err := c.chapterSvc.PageByQueryFilter(ctx.Context(), bookId, queryFilter, pageFilter, sortFilter)
	if err != nil {
		return nil, err
	}

	return &ds.PageDataResponse{
		PageQueryFilter: *pageFilter,
		Total:           total,
		Items:           lo.Map(items, c.chapterSvc.DecorateItem),
	}, nil
}

// Create a new item
//
//	@Summary		Create
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			bookId	path		int	true	"BookId"
//	@Param			body	body		dto.ChapterForm	true	"ChapterForm"
//	@Success		200		{string}	ChapterID
//	@Router			/v1/posts/books/{book_id}/chapters [post]
func (c *ChapterController) Create(ctx *fiber.Ctx, bookId int, body *dto.ChapterForm) error {
	return c.chapterSvc.Create(ctx.Context(), bookId, body)
}

// Update by id
//
//	@Summary		update by id
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			bookId	path		int	true	"BookId"
//	@Param			id		path		int				true	"ChapterID"
//	@Param			body	body		dto.ChapterForm	true	"ChapterForm"
//	@Success		200		{string}	ChapterID
//	@Router			/v1/posts/books/{book_id}/chapters/{id} [put]
func (c *ChapterController) Update(ctx *fiber.Ctx, bookId int, id uint64, body *dto.ChapterForm) error {
	return c.chapterSvc.Update(ctx.Context(), bookId, id, body)
}

// Delete by id
//
//	@Summary		Delete
//	@Tags			DEFAULT_TAG_NAME
//	@Accept			json
//	@Produce		json
//	@Param			bookId	path		int	true	"BookId"
//	@Param			id	path		int	true	"ChapterID"
//	@Success		200	{string}	ChapterID
//	@Router			/v1/posts/books/{book_id}/chapters/{id} [delete]
func (c *ChapterController) Delete(ctx *fiber.Ctx, bookId int, id uint64) error {
	return c.chapterSvc.Delete(ctx.Context(), bookId, id)
}
