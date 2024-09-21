package main

import (
	"cash_register/internal/infrastructure/database"
	"log"
)

func main() {
	//iniciar la conexion a la base de datos
	database.InitDB()

	db := database.GetDB()

	// cerrar la conexion a la base de datos
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error cerrando la conexion a la base de datos: %v", err)
		}
		log.Println("Conexion a la base de datos cerrada")
	}()

}
