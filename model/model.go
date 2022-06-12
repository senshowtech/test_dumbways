package model

type ResponseData struct {
	Status string  `json: "transaction"`
	Price  float32 `json: "price"`
}

type Transaction struct {
	Status string  `json: "transaction"`
	Price  float32 `json: "price"`
}
