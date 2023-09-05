package seeders

import (
	"encoding/json"
	"strings"

	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/docs"
	"github.com/atom-apps/door/modules/auth/dto"
	"github.com/samber/lo"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type RouteSeeder struct {
	id uint64
}

func NewRouteSeeder() contracts.Seeder {
	return &RouteSeeder{}
}

func (s *RouteSeeder) getID() uint64 {
	s.id++
	return s.id
}

func (s *RouteSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Route{}).TableName(nil))

	var doc *dto.SwaggerDoc
	err := json.Unmarshal([]byte(docs.SwaggerSpec), &doc)
	if err != nil {
		return
	}
	ignores := []string{
		"/auth/",
		"/v1/auth/",
		"/v1/services/",
		"/v1/tools/",
		"/v1/users/tokens",
	}

	routes := []models.Route{
		{ID: s.getID(), Type: consts.RouteTypeApi, Name: "用户管理", Method: "ANY", Path: "/v1/users/"},
		{ID: s.getID(), Type: consts.RouteTypeApi, Name: "系统管理", Method: "ANY", Path: "/v1/systems/"},

		{ID: s.getID(), Type: consts.RouteTypeApi, Name: "用户", Method: "ANY", Path: "/v1/users/users", ParentID: 1},
		{ID: s.getID(), Type: consts.RouteTypeApi, Name: "租户", Method: "ANY", Path: "/v1/users/tenants", ParentID: 1},
		{ID: s.getID(), Type: consts.RouteTypeApi, Name: "角色", Method: "ANY", Path: "/v1/users/roles", ParentID: 1},
		{ID: s.getID(), Type: consts.RouteTypeApi, Name: "在线", Method: "ANY", Path: "/v1/users/sessions", ParentID: 1},
		{ID: s.getID(), Type: consts.RouteTypeApi, Name: "权限", Method: "ANY", Path: "/v1/users/permissions", ParentID: 1},

		{ID: s.getID(), Type: consts.RouteTypeApi, Name: "路由", Method: "ANY", Path: "/v1/systems/routes", ParentID: 2},
	}

	genRoutes := lo.FilterMap(doc.ToRoues(), func(item *dto.Route, _ int) (models.Route, bool) {
		ignore := false
		lo.ForEach(ignores, func(i string, _ int) {
			if strings.HasPrefix(item.Path, i) {
				ignore = true
			}
		})

		if ignore {
			return models.Route{}, false
		}

		parentID := uint64(0)
		for _, route := range routes {
			if route.ParentID > 0 && strings.HasPrefix(item.Path, route.Path) {
				parentID = route.ID
			}
		}

		return models.Route{Type: consts.RouteTypeApi, Name: item.Summary, Method: item.Method, Path: item.Path, ParentID: parentID}, true
	})

	routes = append(routes, genRoutes...)
	db.CreateInBatches(&routes, 100)
}

func (s *RouteSeeder) Generate(faker *gofakeit.Faker, idx int) models.Route {
	return models.Route{
		// fill model fields
	}
}
