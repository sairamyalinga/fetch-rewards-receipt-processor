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
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "invalid data format"})
		return
	}

	receipt.ID = uuid.New()
	rp.receipts = append(rp.receipts, receipt)

	ctx.JSON(http.StatusOK, gin.H{"id": receipt.ID})
}
