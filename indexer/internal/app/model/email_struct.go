package model

import "encoding/json"

type JsonParser interface {
	ToJson() ([]byte, error)
}

type BulkV2 struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}

type Email struct {
	MessageID string   `json:"message-id"`
	Date      string   `json:"date"`
	From      string   `json:"from"`
	To        []string `json:"to"`
	Subject   string   `json:"subject"`
	CC        []string `json:"cc"`
	BCC       []string `json:"bcc"`
	XFrom     string   `json:"x-from"`
	XTo       []string `json:"x-to"`
	XCC       []string `json:"x-cc"`
	XBCC      []string `json:"x-bcc"`
	XFolder   string   `json:"x-folder"`
	XFileName string   `json:"x-filename"`
	Content   string   `json:"content"`
}

func (e *BulkV2) ToJson() ([]byte, error) {
	return json.Marshal(e)
}
