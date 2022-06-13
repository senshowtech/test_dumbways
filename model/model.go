package model

type Transaction struct {
	Id     string  `json: "id"`
	Price  float32 `json: "price"`
	Status string  `json: "status"`
}

type Balance struct {
	Id     string  `json: "id"`
	Wallet float32 `json: "wallet"`
	Status string  `json: "status"`
}
