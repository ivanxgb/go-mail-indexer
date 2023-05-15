package models

import "encoding/json"

type ZSearchReq struct {
	SearchType string   `json:"search_type"`
	Source     []string `json:"_source"`
	Query      ZQuery   `json:"query"`
	From       int      `json:"from,omitempty"`
	MaxResult  int      `json:"max_result,omitempty"`
	Size       int      `json:"size,omitempty"`
	SortFields []string `json:"sort_fields,omitempty"`
}

type ZQuery struct {
	Term      string   `json:"term"`
	Field     string   `json:"field,omitempty"`
	Terms     []string `json:"terms,omitempty"`
	StartTime string   `json:"start_time,omitempty"`
	EndTime   string   `json:"end_time,omitempty"`
}

func (e *ZSearchReq) ToJson() ([]byte, error) {
	return json.Marshal(e)
}
