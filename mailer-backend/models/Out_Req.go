package models

type BaseOutReq interface {
	ToJson() ([]byte, error)
}

type Z_Search struct {
	Query Query `json:"query"`
}

type Query struct {
	Sql  string `json:"sql"`
	Size int    `json:"size"`
}

func (e *Z_Search) ToJson() ([]byte, error) {
	return nil, nil
}
