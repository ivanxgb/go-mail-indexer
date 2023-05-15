package models

import "encoding/json"

type SearchResponseOut struct {
	Total int    `json:"total"`
	Mails []Mail `json:"mails"`
}

type Mail struct {
	Id   string `json:"id"`
	Mail Source `json:"mail"`
}

func (e *Hits) ToSearchResponseJson() ([]byte, error) {
	var resp SearchResponseOut
	resp.Total = len(e.Hits)
	resp.Mails = make([]Mail, resp.Total)
	for i, hit := range e.Hits {
		resp.Mails[i].Id = hit.Id
		resp.Mails[i].Mail = hit.Source
	}

	return json.Marshal(resp)
}
