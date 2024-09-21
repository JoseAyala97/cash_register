package models

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	Id             int          `json:"id" gorm:"primaryKey"`
	DenominationId int          `json:"denominationId" gorm:"not null"`
	Denomination   Denomination `json:"denomination" gorm:"foreignKey:DenominationId"`
	Quantity       int          `json:"quantity" gorm:"not null"`
	TotalAmount    float64      `json:"totalAmount" gorm:"not null"`
	TransactionId  int          `json:"transactionId" gorm:"not null"`
	Transaction    Transaction  `json:"transaction" gorm:"foreignKey:TransactionId"`
}
