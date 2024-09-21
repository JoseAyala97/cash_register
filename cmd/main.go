package main

import (
	"cash_register/internal/domain/models"
	"cash_register/internal/infrastructure/database"
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

}
