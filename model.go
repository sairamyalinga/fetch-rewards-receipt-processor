package main

import (
	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}

type Receipt struct {
	Retailer     string    `json:"retailer"`
	PurchaseDate string    `json:"purchaseDate"`
	PurchaseTime string    `json:"purchaseTime"`
	Items        []Item    `json:"items"`
	ID           uuid.UUID `json:"id"`
}

type ReceiptProcessor struct {
	receipts []Receipt
}
