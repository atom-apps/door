// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	 "strings"

	"github.com/atom-apps/door/modules/user/controller"
	"github.com/atom-apps/door/modules/user/dto"
	"github.com/atom-apps/door/common"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeTokenController(engine fiber.Router, controller *controller.TokenController) {
	basePath := "/"+engine.(*fiber.Group).Prefix
	engine.Get(strings.TrimPrefix("/users/tokens/:id<int>", basePath), DataFunc1(controller.Show, Integer[int64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/users/tokens", basePath), DataFunc3(controller.List, Query[dto.TokenListQueryFilter](QueryParamError), Query[common.PageQueryFilter](QueryParamError), Query[common.SortQueryFilter](QueryParamError)))
	engine.Post(strings.TrimPrefix("/users/tokens", basePath), Func1(controller.Create, Body[dto.TokenForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/users/tokens/:id<int>", basePath), Func2(controller.Update, Integer[int64]("id", PathParamError), Body[dto.TokenForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/users/tokens/:id<int>", basePath), Func1(controller.Delete, Integer[int64]("id", PathParamError)))
}
