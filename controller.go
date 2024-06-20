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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid data format"})
		return
	}

	receiptID := uuid.New()
	rp.receipts[receiptID] = receipt

	ctx.JSON(http.StatusOK, gin.H{"id": receiptID})
}

func (rp *ReceiptProcessor) GetPoints(ctx *gin.Context) {
	receiptID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		fmt.Printf("invalid receipt ID: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid receipt ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"points": rp.calculatePoints(receiptID)})
}