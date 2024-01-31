package controllers

import (
	"net/http"

	"enigma.com/m/initializers"
	"enigma.com/m/models"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var body struct {
		CustomerID     int
		PaymentNominal int64
		PaymentType    string
		MerchantID     int
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	history := models.TransactionHistory{
		CustomerID:     body.CustomerID,
		PaymentNominal: body.PaymentNominal,
		PaymentType:    body.PaymentType,
		MerchantID:     body.MerchantID,
	}
	result := initializers.DB.Create(&history)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create transaction",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Payment Successful",
	})

}

func GetTransactionHistory(c *gin.Context) {
	var transactions []models.TransactionHistory

	result := initializers.DB.Find(&transactions)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch transaction list",
		})

		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "History is Empty",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    transactions,
		"message": "Fetch Successful",
	})
}
