package model

type Transaction struct {
	ID     int `json:"id"`
	Type   int `json:"type"`
	UserId int `json:"userId"`
	Value  int `json:"value"`
	Refund int `json:"refund"`
}
