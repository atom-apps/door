package seeders

import (
	"os"

	"github.com/atom-apps/door/common/model"
	"github.com/atom-apps/door/database/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"github.com/rogeecn/fabfile"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
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

type routeDefinition struct {
	Whitelist     []string `json:"whitelist,omitempty"`
	Authorization []*routeItem
}
type routeItem struct {
	Title    string              `json:"name"`
	Name     string              `json:"name"`
	Path     string              `json:"path"`
	Api      []string            `json:"api"`
	Meta     model.RouteMetadata `json:"meta"`
	Order    int                 `json:"order"`
	Children []*routeItem        `json:"children,omitempty"`
}

func (s *RouteSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Route{}).TableName(nil))
	dbUtil.TruncateTable(db, (&models.RouteWhitelist{}).TableName(nil))

	var routes []*models.Route
	yamlSpec, err := os.ReadFile(fabfile.MustFind("routes.yaml"))
	if err != nil {
		panic(err)
	}
	var routesSpec *routeDefinition
	if err := yaml.Unmarshal(yamlSpec, &routesSpec); err != nil {
		panic(err)
	}

	routes = s.Generate(routesSpec.Authorization, 0, "/")
	db.CreateInBatches(&routes, 100)

	whitelistRoutes := lo.Map(routesSpec.Whitelist, func(item string, _ int) *models.RouteWhitelist {
		return &models.RouteWhitelist{
			Route: item,
		}
	})
	db.CreateInBatches(&whitelistRoutes, 100)
}

func (s *RouteSeeder) Generate(items []*routeItem, parentID uint64, prefix string) []*models.Route {
	routes := []*models.Route{}
	for _, item := range items {
		// path := strings.Join([]string{strings.Trim(prefix, "/"), strings.Trim(item.Path, "/")}, "/")
		// path = "/" + strings.TrimLeft(path, "/")

		if item.Meta.Title == "" {
			item.Meta.Title = item.Title
		}

		if item.Meta.RequiresAuth == nil {
			item.Meta.RequiresAuth = lo.ToPtr(true)
		}

		if item.Meta.HideInMenu == nil {
			item.Meta.HideInMenu = lo.ToPtr(true)
		}

		item.Meta.Order = item.Order

		api := []string{}
		if len(item.Api) > 0 {
			api = append(api, item.Api...)
		}

		route := &models.Route{
			ID:       s.getID(),
			Name:     item.Name,
			Path:     item.Path,
			ParentID: parentID,
			Metadata: item.Meta,
			API:      api,
		}
		routes = append(routes, route)

		if len(item.Children) > 0 {
			routes = append(routes, s.Generate(item.Children, route.ID, route.Path)...)
		}
	}
	return routes
}
