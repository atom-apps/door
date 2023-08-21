// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	"strings"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/modules/user/controller"
	"github.com/atom-apps/door/modules/user/dto"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeSessionController(engine fiber.Router, controller *controller.SessionController) {
	basePath := "/" + engine.(*fiber.Group).Prefix
	engine.Get(strings.TrimPrefix("/v1/users/sessions/:id<int>", basePath), DataFunc1(controller.Show, Integer[int64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/v1/users/sessions", basePath), DataFunc3(controller.List, Query[dto.SessionListQueryFilter](QueryParamError), Query[common.PageQueryFilter](QueryParamError), Query[common.SortQueryFilter](QueryParamError)))
	engine.Post(strings.TrimPrefix("/v1/users/sessions", basePath), Func1(controller.Create, Body[dto.SessionForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/v1/users/sessions/:id<int>", basePath), Func1(controller.Delete, Integer[int64]("id", PathParamError)))
	engine.Delete(strings.TrimPrefix("/v1/users/sessions/:sessId<int>/by-session-id", basePath), Func1(controller.DeleteBySessionID, String("sessID", PathParamError)))
}
