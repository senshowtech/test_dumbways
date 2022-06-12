package model

type Wallet struct {
	IdUser       float32 `json: "user"`
	Wallet       float32 `json: "price"`
	Pembelian    float32 `json: "pembelian"`
	Statuswallet string  `json: "transaction"`
}
