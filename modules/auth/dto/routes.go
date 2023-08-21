package dto

import "strings"

type SwaggerDoc struct {
	Paths map[string]map[string]routeDefinition
}

type routeDefinition struct {
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Tags        []string `json:"tags"`
}

type Route struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func (doc *SwaggerDoc) ToRoues() []*Route {
	routes := make([]*Route, 0)
	for path, methods := range doc.Paths {
		for method, definition := range methods {
			route := &Route{
				Path:        path,
				Method:      strings.ToUpper(method),
				Summary:     definition.Summary,
				Description: definition.Description,
			}
			routes = append(routes, route)
		}
	}
	return routes
}
