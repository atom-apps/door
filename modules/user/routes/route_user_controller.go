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

func routeUserController(engine fiber.Router, controller *controller.UserController) {
	basePath := "/"+engine.(*fiber.Group).Prefix
	engine.Get(strings.TrimPrefix("/users/:id<int>", basePath), DataFunc1(controller.Show, Integer[int64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/users", basePath), DataFunc3(controller.List, Query[dto.UserListQueryFilter](QueryParamError), Query[common.PageQueryFilter](QueryParamError), Query[common.SortQueryFilter](QueryParamError)))
	engine.Get(strings.TrimPrefix("/users/roles/:roleId", basePath), DataFunc4(controller.Role, Integer[int64]("roleID", PathParamError), Query[dto.UserListQueryFilter](QueryParamError), Query[common.PageQueryFilter](QueryParamError), Query[common.SortQueryFilter](QueryParamError)))
	engine.Get(strings.TrimPrefix("/users/tenants/:tenantId", basePath), DataFunc4(controller.Tenant, Integer[int64]("tenantID", PathParamError), Query[dto.UserListQueryFilter](QueryParamError), Query[common.PageQueryFilter](QueryParamError), Query[common.SortQueryFilter](QueryParamError)))
	engine.Post(strings.TrimPrefix("/users", basePath), Func1(controller.Create, Body[dto.UserForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/users/:id<int>", basePath), Func2(controller.Update, Integer[int64]("id", PathParamError), Body[dto.UserForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/users/:id<int>", basePath), Func1(controller.Delete, Integer[int64]("id", PathParamError)))
}
