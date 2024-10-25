package main

import (
	"asset-management-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/createAsset", handlers.CreateAsset)
	router.PUT("/updateAsset", handlers.UpdateAsset)
	router.GET("/queryAsset/:dealerID", handlers.QueryAsset)
	router.GET("/transactionHistory/:dealerID", handlers.GetTransactionHistory)

	router.Run(":8080")
}
