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
func (h *TransactionHandler) MakePayment(c *gin.Context) {
	var paymentDTO dtos.PaymentDTO

	// Parsear la solicitud
	if err := c.ShouldBindJSON(&paymentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Procesar el pago y el cambio
	err := h.transactionRegisterUsecase.MakePayment(c.Request.Context(), paymentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusCreated, gin.H{"message": "Pago registrado exitosamente, cambio devuelto"})
}
