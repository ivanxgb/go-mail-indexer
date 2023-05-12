package models

import "encoding/json"

type BaseOutReq interface {
	ToJson() ([]byte, error)
	ToJSONString() (string, error)
}

type ZSearch struct {
	SearchType string   `json:"search_type"`
	Query      ZQuery   `json:"query"`
	SortFields []string `json:"sort_fields"`
	Size       int      `json:"size"`
	Source     []string `json:"_source"`
}

type ZQuery struct {
	Term  string `json:"term"`
	Field string `json:"field"`
}

func (e *ZSearch) ToJson() ([]byte, error) {
	return json.Marshal(e)
}

func (e *ZSearch) ToJSONString() string {
	body, _ := e.ToJson()
	return string(body)
}
