package seeders

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/atom-apps/door/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"github.com/rogeecn/fabfile"
	"gorm.io/datatypes"
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
	Name     string       `json:"name,omitempty"`
	Path     string       `json:"path,omitempty"`
	Api      []string     `json:"api,omitempty"`
	Children []*routeItem `json:"children,omitempty"`
}

func (s *RouteSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Route{}).TableName(nil))

	var routes []*models.Route
	jsonSpec, err := os.ReadFile(fabfile.MustFind("routes.json"))
	if err != nil {
		panic(err)
	}
	var routesSpec *routeDefinition
	if err := json.Unmarshal(jsonSpec, &routesSpec); err != nil {
		panic(err)
	}

	routes = s.Generate(routesSpec.Authorization, 0, "/")

	db.CreateInBatches(&routes, 100)
}

func (s *RouteSeeder) Generate(items []*routeItem, parentID uint64, prefix string) []*models.Route {
	routes := []*models.Route{}
	for _, item := range items {
		path := strings.Join([]string{strings.Trim(prefix, "/"), strings.Trim(item.Path, "/")}, "/")
		path = "/" + strings.TrimLeft(path, "/")

		api := []string{}
		if len(item.Api) > 0 {
			api = append(api, item.Api...)
		}

		route := &models.Route{
			ID:       s.getID(),
			Name:     item.Name,
			Path:     path,
			ParentID: parentID,
			API:      datatypes.JSONType[[]string]{Data: api},
		}
		routes = append(routes, route)

		if len(item.Children) > 0 {
			routes = append(routes, s.Generate(item.Children, route.ID, route.Path)...)
		}
	}
	return routes
}
