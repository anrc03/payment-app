package controllers

import (
	"net/http"

	"enigma.com/m/initializers"
	"enigma.com/m/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterCustomer(c *gin.Context) {
	var body struct {
		FirstName   string
		LastName    string
		Address     string
		PhoneNumber string
		User        models.User
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.User.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	customer := models.Customer{
		FirstName:   body.FirstName,
		LastName:    body.LastName,
		Address:     body.Address,
		PhoneNumber: body.PhoneNumber,
		User: models.User{Email: body.User.Email, Password: string(hash)},
	}
	result := initializers.DB.Create(&customer)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create customer",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully Registered",
	})

}

func GetAllCustomer(c *gin.Context) {
	var customer []models.Customer

	result := initializers.DB.Find(&customer)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch customer list",
		})

		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No Customer Registered",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    customer,
		"message": "Fetch Successful",
	})
}
