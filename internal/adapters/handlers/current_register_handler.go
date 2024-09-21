package handlers

import (
	"cash_register/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CurrentRegisterHandler struct {
	currentRegisterUsecase *usecases.CurrentRegisterUsecase
}

// NewCurrentRegisterHandler es el constructor del handler
func NewCurrentRegisterHandler(cru *usecases.CurrentRegisterUsecase) *CurrentRegisterHandler {
	return &CurrentRegisterHandler{
		currentRegisterUsecase: cru,
	}
}

// Obtener el estado actual de la caja
func (h *CurrentRegisterHandler) GetCurrentRegister(c *gin.Context) {
	currentRegisterViews, totalGeneral, err := h.currentRegisterUsecase.GetCurrentRegister()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el estado de la caja"})
		return
	}

	// Respuesta con las denominaciones y el total general
	c.JSON(http.StatusOK, gin.H{
		"currentRegister": currentRegisterViews,
		"totalGeneral":    totalGeneral,
	})
}
