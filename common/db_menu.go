package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type MenuMetadata struct {
	Icon  string
	Route string
}

func (j MenuMetadata) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan scan value into JSONType[T], implements sql.Scanner interface
func (j *MenuMetadata) Scan(value interface{}) error {
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
