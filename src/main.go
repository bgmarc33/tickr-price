package main

import (
	"github.com/gin-gonic/gin"
	"stockticker/src/service"
)

func main() {
	r := gin.New()
	r.GET("/stock/:ticker", getStock)

	_ = r.Run(":8080")
}

func getStock(ctx *gin.Context) {
	ticker := ctx.Param("ticker")
	if len(ticker) > 8 || len(ticker) < 1 {
		ctx.JSON(422, gin.H{"message": "invalid ticker", "error": true})
		return
	}

	price, err := service.LoadStockPrice(ticker)
	if err != nil {
		ctx.JSON(520, gin.H{"message": err.Error(), "error": true})
		return
	}

	ctx.JSON(200, gin.H{"error": false, "price": price})
}


