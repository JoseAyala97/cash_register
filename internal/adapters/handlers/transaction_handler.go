package handlers

import (
	"cash_register/internal/adapters/dtos"
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

	// Enviar la transacción al caso de uso para procesarla
	err := h.transactionRegisterUsecase.RegisterTransaction(c.Request.Context(), transactionDTO)
	if err != nil {
		// Enviar una respuesta HTTP con el error si ocurre un fallo al registrar
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver una respuesta exitosa
	c.JSON(http.StatusCreated, gin.H{"message": "Transacción registrada exitosamente"})
}
