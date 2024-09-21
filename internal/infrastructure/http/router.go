package http

import (
	"cash_register/internal/adapters/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(denominationHandler *handlers.DenominationHandler, moneyTypeHandler *handlers.MoneyTypeHandler) *gin.Engine {
	router := gin.Default()

	// Rutas para denominaciones
	router.GET("/denominations", denominationHandler.GetAllDenominations)
	router.POST("/denominations", denominationHandler.CreateDenomination)
	router.GET("/denominations/:id", denominationHandler.GetDenominationByID)
	router.PUT("/denominations/:id", denominationHandler.UpdateDenomination)
	router.DELETE("/denominations/:id", denominationHandler.DeleteDenomination)

	// Rutas para tipos de moneda
	router.GET("/money_types", moneyTypeHandler.GetAllMoneyTypes)
	router.POST("/money_types", moneyTypeHandler.CreateMoneyType)
	router.GET("/money_types/:id", moneyTypeHandler.GetMoneyTypeByID)
	router.PUT("/money_types/:id", moneyTypeHandler.UpdateMoneyType)
	router.DELETE("/money_types/:id", moneyTypeHandler.DeleteMoneyType)

	return router
}
