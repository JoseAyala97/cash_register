package handlers

import (
	"cash_register/internal/usecases"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionDetailHandler struct {
	transactionDetailUsecase *usecases.TransactionDetailUsecase
}

// Constructor para el handler
func NewTransactionDetailHandler(tru *usecases.TransactionDetailUsecase) *TransactionDetailHandler {
	return &TransactionDetailHandler{
		transactionDetailUsecase: tru,
	}
}

// Obtener el log de transacciones con filtros opcionales de fecha
func (h *TransactionDetailHandler) GetTransactionLogs(c *gin.Context) {
	// Obtener los parámetros de fecha como cadenas
	startDateStr := c.Query("startDateTime")
	endDateStr := c.Query("endDateTime")

	var startDateTime, endDateTime *time.Time
	var err error

	// Si se proporciona un startDateTime, intenta parsearlo
	if startDateStr != "" {
		parsedStartDateTime, err := time.Parse("2006-01-02T15:04:05", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de startDateTime inválido"})
			return
		}
		startDateTime = &parsedStartDateTime
	}

	// Si se proporciona un endDateTime, intenta parsearlo
	if endDateStr != "" {
		parsedEndDateTime, err := time.Parse("2006-01-02T15:04:05", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de endDateTime inválido"})
			return
		}
		endDateTime = &parsedEndDateTime
	}

	// Llamar al caso de uso para obtener los logs usando las fechas
	logs, err := h.transactionDetailUsecase.GetTransactionLogs(c.Request.Context(), startDateTime, endDateTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el log de transacciones"})
		return
	}

	// Retornar el log de transacciones como respuesta JSON
	c.JSON(http.StatusOK, logs)
}
