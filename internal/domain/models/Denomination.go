package models

import "gorm.io/gorm"

type Denomination struct {
	gorm.Model
	Id          int       `json:"id" gorm:"primaryKey"`
	Value       float64   `json:"value" gorm:"not null"`
	MoneyTypeId int       `json:"moneyTypeId" gorm:"not null"`
	MoneyType   MoneyType `json:"moneyType" gorm:"foreignKey:MoneyTypeId;references:Id"`
}
