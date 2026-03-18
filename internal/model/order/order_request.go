package model

type CreateOrderRequest struct {
	Customer string  `json:"customer"`
	Amount   float64 `json:"amount"`
}
