// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	"strings"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/modules/posts/controller"
	"github.com/atom-apps/door/modules/posts/dto"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeChapterController(engine fiber.Router, controller *controller.ChapterController) {
	groupPrefix := "/" + strings.TrimLeft(engine.(*fiber.Group).Prefix, "/")
	engine.Get(strings.TrimPrefix("/v1/posts/books/:bookId<int>/chapters/:id<int>", groupPrefix), DataFunc2(controller.Show, Integer[uint64]("bookId", PathParamError), Integer[uint64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/v1/posts/books/:bookId<int>/chapters", groupPrefix), DataFunc4(controller.List, Integer[uint64]("bookId", PathParamError), Query[dto.ChapterListQueryFilter](QueryParamError), Query[ds.PageQueryFilter](QueryParamError), Query[ds.SortQueryFilter](QueryParamError)))
	engine.Post(strings.TrimPrefix("/v1/posts/books/:bookId<int>/chapters", groupPrefix), Func2(controller.Create, Integer[int]("bookId", PathParamError), Body[dto.ChapterForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/v1/posts/books/:bookId<int>/chapters/:id<int>", groupPrefix), Func3(controller.Update, Integer[int]("bookId", PathParamError), Integer[uint64]("id", PathParamError), Body[dto.ChapterForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/v1/posts/books/:bookId<int>/chapters/:id<int>", groupPrefix), Func2(controller.Delete, Integer[int]("bookId", PathParamError), Integer[uint64]("id", PathParamError)))
}
