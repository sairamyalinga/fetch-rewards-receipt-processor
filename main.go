package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	rp := ReceiptProcessor{}

	r := gin.Default()
	receiptsRouter := r.Group("/receipts")
	{
		receiptsRouter.POST("/process", rp.AddReceipt)
	}

	r.Run()
}
