package http

import (
	"cash_register/internal/adapters/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(denominationHandler *handlers.DenominationHandler) *gin.Engine {
	router := gin.Default()

	// Rutas para denominaciones
	router.GET("/denominations", denominationHandler.GetAllDenominations)
	router.POST("/denominations", denominationHandler.CreateDenomination)
	router.GET("/denominations/:id", denominationHandler.GetDenominationByID)
	router.PUT("/denominations/:id", denominationHandler.UpdateDenomination)
	router.DELETE("/denominations/:id", denominationHandler.DeleteDenomination)

	return router
}
