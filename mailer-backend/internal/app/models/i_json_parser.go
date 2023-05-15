package models

type jsonParser interface {
	ToJson() ([]byte, error)
}
