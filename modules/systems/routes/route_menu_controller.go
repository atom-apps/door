// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	"strings"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/modules/systems/controller"
	"github.com/atom-apps/door/modules/systems/dto"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeMenuController(engine fiber.Router, controller *controller.MenuController) {
	groupPrefix := "/" + strings.TrimLeft(engine.(*fiber.Group).Prefix, "/")
	engine.Get(strings.TrimPrefix("/v1/systems/menus/:id<int>", groupPrefix), DataFunc1(controller.Show, Integer[uint64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/v1/systems/menus/:id<int>/tree", groupPrefix), DataFunc1(controller.ShowTree, Integer[uint64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/v1/systems/menus", groupPrefix), DataFunc2(controller.List, Query[dto.MenuListQueryFilter](QueryParamError), Query[ds.SortQueryFilter](QueryParamError)))
	engine.Post(strings.TrimPrefix("/v1/systems/menus", groupPrefix), Func1(controller.Create, Body[dto.MenuForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/v1/systems/menus/:menuId<int>/sub", groupPrefix), Func2(controller.CreateSub, Integer[uint64]("menuID", PathParamError), Body[dto.MenuForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/v1/systems/menus/:id<int>", groupPrefix), Func2(controller.Update, Integer[uint64]("id", PathParamError), Body[dto.MenuForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/v1/systems/menus/:id<int>", groupPrefix), Func1(controller.Delete, Integer[uint64]("id", PathParamError)))
}