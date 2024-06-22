package main

import (
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (rp *ReceiptProcessor) calculatePoints(receiptID uuid.UUID) uint64 {
	var points uint64

	receipt := rp.receipts[receiptID]

	// One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points += 1
		}
	}

	// 50 points if the total is a round dollar amount with no cents.
	if receipt.Total == float64(int64(receipt.Total)) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25.
	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}

	// 5 points for every two items on the receipt.
	points += uint64((len(receipt.Items) / 2) * 5)

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	// The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if math.Mod(float64(len(trimmedDescription)), 3) == 0 {
			points += uint64(math.Ceil(item.Price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd.
	// Assumption: ignoring the error returned from .Parse(), assuming the date time input is valid
	receiptDateTime, _ := time.Parse("2006-01-02 15:04", receipt.PurchaseDate+" "+receipt.PurchaseTime)
	if math.Mod(float64(receiptDateTime.Day()), 2) == 1 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	// Assumption: 2:00pm and 4:00pm are inclusive
	if receiptDateTime.Hour() >= 14 && receiptDateTime.Hour() < 16 {
		points += 10
	}

	return points
}
