package models

import "gorm.io/gorm"

type CurrentRegister struct {
	gorm.Model
	Id             int          `json:"id" gorm:"primaryKey"`
	DenominationId int          `json:"denominationId" gorm:"not null"`
	Denomination   Denomination `json:"denomination" gorm:"foreignKey:DenominationId"`
	Quantity       int          `json:"quantity" gorm:"not null"`
}
