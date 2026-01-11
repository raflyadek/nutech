package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	  _ "github.com/lib/pq"

)

func ConnectionDb() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("failed load env %s", err)
	}

	// dbUserame := os.Getenv("DB_USERNAME")
	// dbName := os.Getenv("DB_NAME")
	// pgPort := os.Getenv("PGPORT")
	// host := os.Getenv("HOST")
	postgreURL := os.Getenv("POSTGRE_URL")
	// dsn := fmt.Sprintf("")

	db, err := sql.Open("postgres", postgreURL)
	if err != nil {
		log.Printf("failed open db %s", err)
	}

	// flag if success
	fmt.Println("success connect to db!")
	return db, nil
}