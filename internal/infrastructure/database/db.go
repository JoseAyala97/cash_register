package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

var db *sql.DB

// cargar variables de entorno y establecer la conexion
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

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error abriendo la conexion: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error verificando la conexion a la base de datos: %v", err)
	}
	log.Println("Conectado a la base de datos")
}

// GetDB devuelve la instancia de la base de datos
func GetDB() *sql.DB {
	return db
}
