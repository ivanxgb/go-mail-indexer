package models

import (
	"encoding/json"
	"time"
)

type ZSearchResponse struct {
	Took     int     `json:"took"`
	TimedOut bool    `json:"timed_out"`
	MaxScore float64 `json:"max_score"`
	Hits     Hits    `json:"hits"`
	Error    string  `json:"error"`
}

type Hits struct {
	Total Total `json:"total"`
	Hits  []Hit `json:"hits"`
}

type Total struct {
	Value int `json:"value"`
}

type Hit struct {
	Index     string  `json:"_index"`
	Type      string  `json:"_type"`
	Id        string  `json:"_id"`
	Score     float64 `json:"_score"`
	Timestamp string  `json:"@timestamp"`
	Source    Source  `json:"_source"`
}

type Source struct {
	Timestamp time.Time `json:"@timestamp"`
	Bcc       []any     `json:"bcc"`
	Cc        []any     `json:"cc"`
	Content   string    `json:"content"`
	Date      string    `json:"date"`
	From      string    `json:"from"`
	MessageID string    `json:"message-id"`
	Subject   string    `json:"subject"`
	To        []string  `json:"to"`
	XBcc      []any     `json:"x-bcc"`
	XCc       []any     `json:"x-cc"`
	XFilename string    `json:"x-filename"`
	XFolder   string    `json:"x-folder"`
	XFrom     string    `json:"x-from"`
	XTo       []string  `json:"x-to"`
}

func (e *ZSearchResponse) FromJson(data []byte) error {
	err := json.Unmarshal(data, &e)
	if err != nil {
		return err
	}

	return nil
}

func (e *Hits) ToJson() ([]byte, error) {
	return json.Marshal(e)
}
