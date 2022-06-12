package model

type Transaction struct {
	Id     float32 `json: "id"`
	Price  float32 `json: "price"`
	Status string  `json: "status"`
}
