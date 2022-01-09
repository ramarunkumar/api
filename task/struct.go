package main

type Seller struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phoneno string `json:"phoneno"`
	Role    string `json:"role"`
}

type Buyer struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email" `
	Phoneno int64  `json:"phoneno"`
	Role    string `json:"role"`
}

type Product struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Tax       float64 `json:"tax"`
	Price     float64 `json:"price"`
	Seller_id string  `json:"seller_id"`
}
