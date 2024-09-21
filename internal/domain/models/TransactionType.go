package models

import "gorm.io/gorm"

type TransactionType struct {
	gorm.Model
	Id    int    `json:"id" gorm:"primaryKey"`
	Value string `json:"value" gorm:"not null"`
}
