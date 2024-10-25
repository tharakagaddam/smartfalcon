package handlers

import (
	"asset-management-api/gateway"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAssetRequest struct {
	DealerID string  `json:"dealerID" binding:"required"`
	MSISDN   string  `json:"msisdn" binding:"required"`
	MPIN     string  `json:"mpin" binding:"required"`
	Balance  float64 `json:"balance" binding:"required"`
	Status   string  `json:"status" binding:"required"`
	Remarks  string  `json:"remarks"`
}

type UpdateAssetRequest struct {
	DealerID string  `json:"dealerID" binding:"required"`
	Balance  float64 `json:"balance" binding:"required"`
	Status   string  `json:"status" binding:"required"`
	Remarks  string  `json:"remarks"`
}

func CreateAsset(c *gin.Context) {
	var req CreateAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := gateway.CreateAsset(req.DealerID, req.MSISDN, req.MPIN, req.Balance, req.Status, req.Remarks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset created successfully"})
}

func UpdateAsset(c *gin.Context) {
	var req UpdateAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := gateway.UpdateAsset(req.DealerID, req.Balance, req.Status, req.Remarks)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asset updated successfully"})
}

func QueryAsset(c *gin.Context) {
	dealerID := c.Param("dealerID")
	asset, err := gateway.QueryAsset(dealerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	assetJSON, _ := json.Marshal(asset)
	c.Data(http.StatusOK, "application/json", assetJSON)
}

func GetTransactionHistory(c *gin.Context) {
	dealerID := c.Param("dealerID")
	history, err := gateway.GetTransactionHistory(dealerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	historyJSON, _ := json.Marshal(history)
	c.Data(http.StatusOK, "application/json", historyJSON)
}
