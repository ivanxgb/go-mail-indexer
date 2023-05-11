package models

import (
	"encoding/json"
	"fmt"
	"io"
)

type BaseReq interface {
	FromJson(data io.ReadCloser) (interface{}, error)
}

type Search struct {
	Search string `json:"search"`
}

func (e *Search) FromJson(data io.ReadCloser) error {
	decoder := json.NewDecoder(data)
	err := decoder.Decode(&e)
	if err != nil && err != io.EOF {
		return err
	}

	if e.Search == "" {
		return fmt.Errorf("incorrect json format")
	}

	return nil
}
