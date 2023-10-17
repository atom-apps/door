// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	"strings"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/modules/users/controller"
	"github.com/atom-apps/door/modules/users/dto"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeTokenController(engine fiber.Router, controller *controller.TokenController) {
	groupPrefix := "/" + strings.TrimLeft(engine.(*fiber.Group).Prefix, "/")
	engine.Get(strings.TrimPrefix("/v1/users/tokens", groupPrefix), DataFunc3(controller.List, Query[dto.TokenListQueryFilter](QueryParamError), Query[ds.PageQueryFilter](QueryParamError), Query[ds.SortQueryFilter](QueryParamError)))
	engine.Delete(strings.TrimPrefix("/v1/users/tokens/:id<int>", groupPrefix), Func1(controller.Delete, Integer[uint64]("id", PathParamError)))
}
