package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	FirstName string
	LastName string
	Address string
	PhoneNumber string
	UserID int
	User User `gorm:"references:id"`
}