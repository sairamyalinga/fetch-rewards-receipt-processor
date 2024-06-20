package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (rp *ReceiptProcessor) AddReceipt(ctx *gin.Context) {
	var receipt Receipt

	if err := ctx.ShouldBindJSON(&receipt); err != nil {
		fmt.Printf("invalid data format: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "The receipt is invalid"})
		return
	}

	receiptID := uuid.New()
	rp.receipts[receiptID] = receipt

	ctx.JSON(http.StatusOK, gin.H{"id": receiptID})
}

func (rp *ReceiptProcessor) GetPoints(ctx *gin.Context) {
	receiptID, _ := uuid.Parse(ctx.Param("id"))
	if _, ok := rp.receipts[receiptID]; !ok {
		fmt.Printf("invalid receipt ID: %s\n", ctx.Param("id"))
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No receipt found for that id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"points": rp.calculatePoints(receiptID)})
}