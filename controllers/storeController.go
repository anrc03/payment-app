package controllers

import (
	"net/http"

	"enigma.com/m/initializers"
	"enigma.com/m/models"
	"github.com/gin-gonic/gin"
)

func RegisterStore(c *gin.Context) {
	var body struct {
		NoSiup   string
		StoreName    string
		Address     string
		PhoneNumber string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	store := models.Store{
		NoSiup:   body.NoSiup,
		StoreName:    body.StoreName,
		Address:     body.Address,
		PhoneNumber: body.PhoneNumber,
	}
	result := initializers.DB.Create(&store)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create store",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Store Succesfully Registered",
	})

}

func GetAllStore(c *gin.Context) {
	var store []models.Store

	result := initializers.DB.Find(&store)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch store list",
		})

		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No Store is Registered",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    store,
		"message": "Fetch Successful",
	})
}