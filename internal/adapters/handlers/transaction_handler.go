package handlers

import (
	"cash_register/internal/adapters/dtos"
	"cash_register/internal/domain/models"
	"cash_register/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionRegisterUsecase *usecases.TransactionRegister
}

func NewTransactionHandler(tru *usecases.TransactionRegister) *TransactionHandler {
	return &TransactionHandler{
		transactionRegisterUsecase: tru,
	}
}

// Registrar una transacción con sus detalles
func (h *TransactionHandler) RegisterTransaction(c *gin.Context) {
	var transactionDTO dtos.TransactionDTO

	// Parsear el cuerpo de la solicitud para obtener la transacción y los detalles
	if err := c.ShouldBindJSON(&transactionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Crear la entidad Transaction
	transaction := models.Transaction{
		TransactionTypeId: transactionDTO.TransactionTypeId,
	}

	// Registrar la transacción con detalles (se calculará el TotalAmount en el Usecase)
	err := h.transactionRegisterUsecase.RegisterTransaction(c.Request.Context(), &transaction, transactionDTO.Details)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar la transacción"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transacción registrada exitosamente"})
}
