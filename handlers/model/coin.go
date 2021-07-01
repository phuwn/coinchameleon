package model

type Coin struct {
	ID     string `json:"id" sql:"id"`
	Name   string `json:"name" sql:"name"`
	Symbol string `json:"symbol" sql:"symbol"`
}
