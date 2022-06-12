package model

type Transaction struct {
	Id     float32 `json: "Id"`
	Price  float32 `json: "price"`
	Status string  `json: "status"`
}
