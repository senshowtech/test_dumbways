package model

type Transaction struct {
	Status string  `json: "transaction"`
	Price  float32 `json: "price"`
}
