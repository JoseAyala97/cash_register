package main

import (
	"cash_register/internal/adapters/handlers"
	"cash_register/internal/adapters/repositories"
	"cash_register/internal/domain/models"
	"cash_register/internal/infrastructure/database"
	"cash_register/internal/infrastructure/http"
	"cash_register/internal/usecases"
	"log"
)

func main() {
	//iniciar la conexion a la base de datos
	database.InitDB()

	// obtener la instancia de la base de datos
	db := database.GetDB()

	// Migrar las tablas basadas en los modelos
	err := db.AutoMigrate(
		&models.CurrentRegister{},
		&models.Denomination{},
		&models.MoneyType{},
		&models.Transaction{},
		&models.TransactionDetail{},
		&models.TransactionType{},
	)
	if err != nil {
		log.Fatalf("Error al migrar las tablas: %v", err)
	}

	log.Println("Migración de tablas completada")

	// Crear repositorios
	denominationRepo := repositories.NewDenominationRepository(db)
	moneyTypeRepo := repositories.NewGenericRepository[models.MoneyType](db)
	transactionTypeRepo := repositories.NewTransactionTypeRepository(db)

	transactionRepo := repositories.NewTransactionRepository(db)
	transactionDetailRepo := repositories.NewTransactionDetailRepository(db)
	currentRegisterRepo := repositories.NewCurrentRegisterRepository(db) // Uso correcto

	// Crear los usecases inyectando los repositorios
	denominationUsecase := usecases.NewDenominationUsecase(denominationRepo)
	moneyTypeUsecase := usecases.NewMoneyTypeUsecase(moneyTypeRepo)
	transactionTypeUsecase := usecases.NewTransactionTypeUsecase(transactionTypeRepo)

	// Aquí pasamos el `currentRegisterRepo` en lugar de `denominationRepo` en la posición correcta
	transactionRegisterUsecase := usecases.NewTransactionRegister(transactionRepo, transactionDetailRepo, currentRegisterRepo, denominationRepo)

	// Crear los handlers inyectando los usecases
	denominationHandler := handlers.NewDenominationHandler(denominationUsecase)
	moneyTypeHandler := handlers.NewMoneyTypeHandler(moneyTypeUsecase)
	transactionTypeHandler := handlers.NewTransactionTypeHandler(transactionTypeUsecase)
	transactionHandler := handlers.NewTransactionHandler(transactionRegisterUsecase)
	currentRegisterHandler := handlers.NewCurrentRegisterHandler(usecases.NewCurrentRegisterUsecase(currentRegisterRepo))

	// Configurar las rutas con los handlers
	router := http.SetupRouter(denominationHandler, moneyTypeHandler, transactionTypeHandler, transactionHandler, currentRegisterHandler)

	err = router.Run(":8080") // Aquí inicia el servidor y lo mantiene activo
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

}
