package http

import (
	"cash_register/internal/adapters/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(denominationHandler *handlers.DenominationHandler, moneyTypeHandler *handlers.MoneyTypeHandler, transactionTypeHandler *handlers.TransactionTypeHandler) *gin.Engine {
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

	// Rutas para tipos de transacción
	router.GET("/transaction_types", transactionTypeHandler.GetAllTransactionTypes)
	router.POST("/transaction_types", transactionTypeHandler.CreateTransactionType)
	router.GET("/transaction_types/:id", transactionTypeHandler.GetTransactionTypeByID)
	router.PUT("/transaction_types/:id", transactionTypeHandler.UpdateTransactionType)
	router.DELETE("/transaction_types/:id", transactionTypeHandler.DeleteTransactionType)

	return router
}
