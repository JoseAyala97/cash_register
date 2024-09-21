package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Id                int             `json:"id" gorm:"primaryKey"`
	TransactionTypeId int             `json:"transactionTypeId" gorm:"not null"`
	TransactionType   TransactionType `json:"transactionType" gorm:"foreignKey:TransactionTypeId"`
	TotalAmount       float64         `json:"totalAmount" gorm:"not null"`
	PaidAmount        float64         `json:"paidAmount"`
}
