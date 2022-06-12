package model

type Wallet struct {
	Status string  `json: "transaction"`
	Price  float32 `json: "price"`
}
