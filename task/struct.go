package main

type Users struct {
	Id      string `json:"id" bind:"id"`
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Phoneno string `json:"phoneno" binding:"required"`
	Role    string `json:"role"`
}

type Product struct {
	Id        string `json:"id" `
	Name      string `json:"name" binding:"required"`
	Price     string `json:"price" binding:"required"`
	Tax       string `json:"tax" binding:"required"`
	Seller_id string `json:"seller_id" binding:"required"`
}
type Order struct {
	Id        string `json:"id" `
	Email     string `json:"email" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Price     string `json:"price" binding:"required"`
	Tax       string `json:"tax" binding:"required"`
	Seller_id string `json:"seller_id" binding:"required"`
	Quantity  string `json:"quantity" binding:"required"`
}
