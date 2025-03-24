package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	log.Printf("DB_HOST: %s", dbHost)
	log.Printf("DB_PORT: %s", dbPort)
	log.Printf("DB_USER: %s", dbUser)
	log.Printf("DB_PASSWORD: %s", dbPassword)
	log.Printf("DB_NAME: %s", dbName)

	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPassword,
		dbName,
	)

	db, err := sql.Open("postgres", connectString)
	if err != nil {
		return nil, err
	}

	return db, db.Ping()
}
