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

func routeTenantController(engine fiber.Router, controller *controller.TenantController) {
	basePath := "/"+engine.(*fiber.Group).Prefix
	engine.Get(strings.TrimPrefix("/user/tenants/:id<int>", basePath), DataFunc1(controller.Show, Integer[int64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/user/tenants", basePath), DataFunc3(controller.List, Query[dto.TenantListQueryFilter](QueryParamError), Query[common.PageQueryFilter](QueryParamError), Query[common.SortQueryFilter](QueryParamError)))
	engine.Post(strings.TrimPrefix("/user/tenants", basePath), Func1(controller.Create, Body[dto.TenantForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/user/tenants/:id<int>", basePath), Func2(controller.Update, Integer[int64]("id", PathParamError), Body[dto.TenantForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/user/tenants/:id<int>", basePath), Func1(controller.Delete, Integer[int64]("id", PathParamError)))
}
