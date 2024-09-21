package handlers

import (
	"cash_register/internal/domain/models"
	"cash_register/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionTypeHandler struct {
	usecase *usecases.TransactionTypeUsecase
}

// NewTransactionTypeHandler crea un nuevo handler para TransactionType
func NewTransactionTypeHandler(usecase *usecases.TransactionTypeUsecase) *TransactionTypeHandler {
	return &TransactionTypeHandler{usecase: usecase}
}

// Obtener todos los tipos de transacciones
func (h *TransactionTypeHandler) GetAllTransactionTypes(c *gin.Context) {
	types, err := h.usecase.GetAllTransactionTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los tipos de transacciones"})
		return
	}
	c.JSON(http.StatusOK, types)
}

// Crear un nuevo tipo de transacción
func (h *TransactionTypeHandler) CreateTransactionType(c *gin.Context) {
	var transactionType models.TransactionType
	if err := c.ShouldBindJSON(&transactionType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err := h.usecase.CreateTransactionType(transactionType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el tipo de transacción"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tipo de transacción creado"})
}

// Obtener un tipo de transacción por ID
func (h *TransactionTypeHandler) GetTransactionTypeByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	transactionType, err := h.usecase.GetTransactionTypeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tipo de transacción no encontrado"})
		return
	}

	c.JSON(http.StatusOK, transactionType)
}

// Actualizar un tipo de transacción
func (h *TransactionTypeHandler) UpdateTransactionType(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var transactionType models.TransactionType
	if err := c.ShouldBindJSON(&transactionType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	err = h.usecase.UpdateTransactionType(uint(id), transactionType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el tipo de transacción"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tipo de transacción actualizado"})
}

// Eliminar un tipo de transacción
func (h *TransactionTypeHandler) DeleteTransactionType(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.usecase.DeleteTransactionType(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el tipo de transacción"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tipo de transacción eliminado"})
}
