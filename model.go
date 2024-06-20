package main

import (
	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total,string"`
}

type ReceiptProcessor struct {
	receipts map[uuid.UUID]Receipt
}
