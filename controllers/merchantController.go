package controllers

import (
	"net/http"

	"enigma.com/m/initializers"
	"enigma.com/m/models"
	"github.com/gin-gonic/gin"
)

func RegisterMerchant(c *gin.Context) {
	var body struct {
		FullName   string
		Address     string
		PhoneNumber string
		Store models.Store
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	merchant := models.Merchant{
		FullName:   body.FullName,
		Address:     body.Address,
		PhoneNumber: body.PhoneNumber,
		Store: models.Store{
			NoSiup:   body.Store.NoSiup,
			StoreName:    body.Store.StoreName,
			Address:     body.Store.Address,
			PhoneNumber: body.Store.PhoneNumber,
		},
	}
	result := initializers.DB.Create(&merchant)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create merchant account and store",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully Registered",
	})

}

func GetAllMerchant(c *gin.Context) {
	var merchant []models.Merchant

	result := initializers.DB.Find(&merchant)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch merchant list",
		})

		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No Merchant(s) Registered",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    merchant,
		"message": "Fetch Successful",
	})
}