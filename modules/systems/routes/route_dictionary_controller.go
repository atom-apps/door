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

func routeDictionaryController(engine fiber.Router, controller *controller.DictionaryController) {
	groupPrefix := "/" + strings.TrimLeft(engine.(*fiber.Group).Prefix, "/")
	engine.Get(strings.TrimPrefix("/v1/systems/dictionaries/:id<int>", groupPrefix), DataFunc1(controller.Show, Integer[uint64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/v1/systems/dictionaries", groupPrefix), DataFunc3(controller.List, Query[dto.DictionaryListQueryFilter](QueryParamError), Query[ds.PageQueryFilter](QueryParamError), Query[ds.SortQueryFilter](QueryParamError)))
	engine.Post(strings.TrimPrefix("/v1/systems/dictionaries", groupPrefix), Func1(controller.Create, Body[dto.DictionaryForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/v1/systems/dictionaries/:id<int>", groupPrefix), Func2(controller.Update, Integer[uint64]("id", PathParamError), Body[dto.DictionaryForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/v1/systems/dictionaries/:id<int>", groupPrefix), Func1(controller.Delete, Integer[uint64]("id", PathParamError)))
}
