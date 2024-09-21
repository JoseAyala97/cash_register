package handlers

import (
	"cash_register/internal/domain/models"
	"cash_register/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DenominationHandler estructura para manejar las dependencias del usecase
type DenominationHandler struct {
	usecase *usecases.DenominationUsecase
}

// Constructor para crear un DenominationHandler con el usecase inyectado
func NewDenominationHandler(usecase *usecases.DenominationUsecase) *DenominationHandler {
	return &DenominationHandler{usecase: usecase}
}

// Obtener todas las denominaciones
func (h *DenominationHandler) GetAllDenominations(c *gin.Context) {
	denominations, err := h.usecase.GetAllDenominations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener denominaciones"})
		return
	}
	c.JSON(http.StatusOK, denominations)
}

// Crear una nueva denominación
func (h *DenominationHandler) CreateDenomination(c *gin.Context) {
	var denomination models.Denomination
	if err := c.ShouldBindJSON(&denomination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := h.usecase.CreateDenomination(denomination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear denominación"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Denominación creada"})
}

// Obtener una denominación por ID
func (h *DenominationHandler) GetDenominationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	denomination, err := h.usecase.GetDenominationByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Denominación no encontrada"})
		return
	}

	c.JSON(http.StatusOK, denomination)
}

// Actualizar una denominación por ID
func (h *DenominationHandler) UpdateDenomination(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var denomination models.Denomination
	if err := c.ShouldBindJSON(&denomination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err = h.usecase.UpdateDenomination(uint(id), denomination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar denominación"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Denominación actualizada"})
}

// Eliminar una denominación por ID
func (h *DenominationHandler) DeleteDenomination(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.usecase.DeleteDenomination(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar denominación"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Denominación eliminada"})
}
