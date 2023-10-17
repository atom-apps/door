package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/atom-apps/door/common/consts"
)

type ArticleThumbnail struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Image  string `json:"image"`
	Head   bool   `json:"head"`
}

type ArticleThumbnails []ArticleThumbnail

func (j ArticleThumbnails) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *ArticleThumbnails) Scan(value interface{}) error {
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

type ArticleVideo struct {
	Provider consts.AudioProvider `json:"provider"`
	Url      string               `json:"url"`
}

type ArticleVideos []ArticleVideo

func (j ArticleVideos) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *ArticleVideos) Scan(value interface{}) error {
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

type ArticleAudio struct {
	Provider consts.AudioProvider `json:"provider"`
	Url      string               `json:"url"`
}

type ArticleAudios []ArticleAudio

func (j ArticleAudios) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *ArticleAudios) Scan(value interface{}) error {
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
