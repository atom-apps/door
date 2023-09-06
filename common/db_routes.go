package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type RouteAPI []string

func (j RouteAPI) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan scan value into JSONType[T], implements sql.Scanner interface
func (j *RouteAPI) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, &j)
}

type RouteMetadata struct {
	Title        string `json:"title"`
	RequiresAuth *bool  `json:"requiresAuth" yaml:"requireAuth"`
	HideInMenu   *bool  `json:"hideInMenu" yaml:"hideInMenu"`
	Icon         string `json:"icon,omitempty"`
	Order        int    `json:"order"`
}

func (j RouteMetadata) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan scan value into JSONType[T], implements sql.Scanner interface
func (j *RouteMetadata) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, &j)
}

type RouteItem struct {
	Name     string        `json:"name"`
	Path     string        `json:"path"`
	Meta     RouteMetadata `json:"meta"`
	Order    int           `json:"-"`
	Children []*RouteItem  `json:"children,omitempty"`
}
