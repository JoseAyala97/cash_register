package handlers

import (
	"cash_register/internal/domain/models"
	"cash_register/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MoneyTypeHandler struct {
	usecases *usecases.MoneyTypeUsecase
}

func NewMoneyTypeHandler(usecases *usecases.MoneyTypeUsecase) *MoneyTypeHandler {
	return &MoneyTypeHandler{usecases: usecases}
}

func (h *MoneyTypeHandler) GetAllMoneyTypes(c *gin.Context) {
	moneyTypes, err := h.usecases.GetAllMoneyTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener tipos de moneda"})
		return
	}
	c.JSON(http.StatusOK, moneyTypes)
}

func (h *MoneyTypeHandler) CreateMoneyType(c *gin.Context) {
	var moneyType models.MoneyType
	if err := c.ShouldBindJSON(&moneyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := h.usecases.CreateMoneyType(moneyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear tipo de moneda"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tipo de moneda creado"})
}

func (h *MoneyTypeHandler) GetMoneyTypeByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	moneyType, err := h.usecases.GetMoneyTypeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tipo de moneda no encontrado"})
		return
	}

	c.JSON(http.StatusOK, moneyType)
}

func (h *MoneyTypeHandler) UpdateMoneyType(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var moneyType models.MoneyType
	if err := c.ShouldBindJSON(&moneyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err = h.usecases.UpdateMoneyType(uint(id), moneyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar tipo de moneda"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tipo de moneda actualizado"})
}

func (h *MoneyTypeHandler) DeleteMoneyType(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.usecases.DeleteMoneyType(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar tipo de moneda"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tipo de moneda eliminado"})
}
