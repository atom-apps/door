// Code generated by the atomctl; DO NOT EDIT.

package routes

import (
	"strings"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/modules/storages/controller"
	"github.com/atom-apps/door/modules/storages/dto"
	"github.com/atom-providers/jwt"

	"github.com/gofiber/fiber/v2"
	. "github.com/rogeecn/fen"
)

func routeFilesystemController(engine fiber.Router, controller *controller.FilesystemController) {
	groupPrefix := "/" + strings.TrimLeft(engine.(*fiber.Group).Prefix, "/")
	engine.Get(strings.TrimPrefix("/v1/storages/filesystems/:id<int>", groupPrefix), DataFunc1(controller.Show, Integer[uint64]("id", PathParamError)))
	engine.Get(strings.TrimPrefix("/v1/storages/filesystems", groupPrefix), DataFunc4(controller.List, JwtClaim[jwt.Claims](ClaimParamError), Query[dto.FilesystemListQueryFilter](QueryParamError), Query[ds.PageQueryFilter](QueryParamError), Query[ds.SortQueryFilter](QueryParamError)))
	engine.Post(strings.TrimPrefix("/v1/storages/filesystems", groupPrefix), Func1(controller.Create, Body[dto.FilesystemForm](BodyParamError)))
	engine.Put(strings.TrimPrefix("/v1/storages/filesystems/:id<int>", groupPrefix), Func2(controller.Update, Integer[uint64]("id", PathParamError), Body[dto.FilesystemForm](BodyParamError)))
	engine.Delete(strings.TrimPrefix("/v1/storages/filesystems/:id<int>", groupPrefix), Func1(controller.Delete, Integer[uint64]("id", PathParamError)))
	engine.Post(strings.TrimPrefix("/v1/storages/filesystems/:id<int>/directory", groupPrefix), Func3(controller.Directory, JwtClaim[jwt.Claims](ClaimParamError), Integer[uint64]("id", PathParamError), Body[dto.CreateSubDirectoryForm](BodyParamError)))
	engine.Get(strings.TrimPrefix("/v1/storages/filesystems/directories/tree", groupPrefix), DataFunc1(controller.DirectoryTree, JwtClaim[jwt.Claims](ClaimParamError)))
	engine.Post(strings.TrimPrefix("/v1/storages/filesystems/:id<int>/move", groupPrefix), Func3(controller.MoveFiles, JwtClaim[jwt.Claims](ClaimParamError), Integer[uint64]("id", PathParamError), Body[ds.IDsForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/v1/storages/filesystems/:id<int>/copy", groupPrefix), Func3(controller.CopyFiles, JwtClaim[jwt.Claims](ClaimParamError), Integer[uint64]("id", PathParamError), Body[ds.IDsForm](BodyParamError)))
	engine.Post(strings.TrimPrefix("/v1/storages/filesystems/get-by-real-names", groupPrefix), DataFunc2(controller.GetByRealNames, JwtClaim[jwt.Claims](ClaimParamError), Body[dto.RealNamesForm](BodyParamError)))
}
