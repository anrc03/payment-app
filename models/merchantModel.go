package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	FullName string
	Address string
	PhoneNumber string
	StoreID int
	Store Store `gorm:"references:id"`
}