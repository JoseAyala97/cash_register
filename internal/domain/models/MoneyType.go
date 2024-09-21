package models

import "gorm.io/gorm"

type MoneyType struct {
	gorm.Model
	Id    int    `json:"id" gorm:"primaryKey"`
	Value string `json:"value" gorm:"not null"`
}
