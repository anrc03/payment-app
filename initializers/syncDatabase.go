package initializers

import "enigma.com/m/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Store{})
	DB.AutoMigrate(&models.Merchant{})
	DB.AutoMigrate(&models.TransactionHistory{})
}