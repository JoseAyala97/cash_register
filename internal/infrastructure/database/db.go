package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB carga las variables de entorno y establece la conexión usando GORM
func InitDB() {
	err := gotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar variables de entorno: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	// Conectando con GORM
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	log.Println("Conectado a la base de datos con GORM")
}

// GetDB devuelve la instancia de la base de datos GORM
func GetDB() *gorm.DB {
	return db
}
