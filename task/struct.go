package main

type Users struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phoneno string `json:"phoneno"`
	Role    string `json:"role"`
}

type Product struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	Tax       string `json:"tax"`
	Seller_id string `json:"seller_id"`
}
type Order struct {
	Id        string `json:"id"`
	Email     string `json:"email" `
	Name      string `json:"name"`
	Price     string `json:"price"`
	Tax       string `json:"tax"`
	Seller_id string `json:"seller_id"`
	Quantity  string `json:"quantity"`
}
