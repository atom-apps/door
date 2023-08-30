package common

import (
	"strings"
)

type SortQueryFilter struct {
	Asc  *string `json:"asc" form:"asc"`
	Desc *string `json:"desc" form:"desc"`
}

func (s *SortQueryFilter) AscFields() []string {
	if s.Asc == nil {
		return nil
	}
	return strings.Split(*s.Asc, ",")
}

func (s *SortQueryFilter) DescFields() []string {
	if s.Desc == nil {
		return nil
	}
	return strings.Split(*s.Desc, ",")
}

type PageDataResponse struct {
	PageQueryFilter `json:",inline"`
	Total           int64       `json:"total"`
	Items           interface{} `json:"items"`
}

type PageQueryFilter struct {
	Page     int `json:"page,omitempty" form:"page"`
	Limit    int `json:"limit,omitempty" form:"limit"`
	Current  int `json:"current,omitempty" form:"current"`
	PageSize int `json:"page_size,omitempty" form:"pageSize"`
}

func (filter *PageQueryFilter) Offset() int {
	return (filter.Page - 1) * filter.Limit
}

func (filter *PageQueryFilter) Format() *PageQueryFilter {
	if (filter.Page == 0 && filter.Current > 0) || (filter.Limit == 0 && filter.PageSize > 0) {
		filter.Page = filter.Current
		filter.PageSize = filter.Limit
	}

	if filter.Page <= 0 {
		filter.Page = 1
	}

	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	if filter.Limit > 50 {
		filter.Limit = 50
	}
	return filter
}

type IDsForm struct {
	IDs []int64 `json:"ids" form:"ids"`
}

type LabelItem struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
