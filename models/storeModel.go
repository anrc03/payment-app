package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	NoSiup string
	StoreName string
	Address string
	PhoneNumber string
}