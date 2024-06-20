package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	rp := ReceiptProcessor{
		receipts: make(map[uuid.UUID]Receipt),
	}

	r := gin.Default()
	receiptsRouter := r.Group("/receipts")
	{
		receiptsRouter.POST("/process", rp.AddReceipt)
		receiptsRouter.GET("/:id/points", rp.GetPoints)
	}

	r.Run()
}
