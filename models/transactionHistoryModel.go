package models

import "gorm.io/gorm"

type TransactionHistory struct {
	gorm.Model
	CustomerID int
	Customer Customer `gorm:"references:id"`
	PaymentNominal int64
	PaymentType string
	MerchantID int
	Merchant Merchant `gorm:"references:id"`
}