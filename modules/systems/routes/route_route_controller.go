// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	 "strings"

	"github.com/atom-apps/door/modules/systems/controller"
	"github.com/atom-providers/jwt"
	"github.com/atom-apps/door/modules/systems/dto"
	"github.com/atom-apps/door/common"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeRouteController(engine fiber.Router, controller *controller.RouteController) {
	basePath := "/"+engine.(*fiber.Group).Prefix
	engine.Get(strings.TrimPrefix("/v1/systems/routes/:id<int>", basePath), DataFunc2(controller.Show, JwtClaim[jwt.Claims](ClaimParamError), Integer[uint64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/v1/systems/routes", basePath), DataFunc4(controller.List, JwtClaim[jwt.Claims](ClaimParamError), Query[dto.RouteListQueryFilter](QueryParamError), Query[common.PageQueryFilter](QueryParamError), Query[common.SortQueryFilter](QueryParamError)))
	engine.Get(strings.TrimPrefix("/v1/systems/routes/pages", basePath), DataFunc2(controller.Pages, JwtClaim[jwt.Claims](ClaimParamError), String("routeType", PathParamError)))
	engine.Post(strings.TrimPrefix("/v1/systems/routes", basePath), Func2(controller.Create, JwtClaim[jwt.Claims](ClaimParamError), Body[dto.RouteForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/v1/systems/routes/:id<int>", basePath), Func3(controller.Update, JwtClaim[jwt.Claims](ClaimParamError), Integer[uint64]("id", PathParamError), Body[dto.RouteForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/v1/systems/routes/:id<int>", basePath), Func2(controller.Delete, JwtClaim[jwt.Claims](ClaimParamError), Integer[uint64]("id", PathParamError)))
}
