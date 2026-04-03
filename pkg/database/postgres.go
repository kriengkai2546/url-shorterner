package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	 _ "github.com/lib/pq"
)

func Connect() *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Cannot open db:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("db not reachable:", err)
	}

	log.Println("DB connected successfully")
  return db
}
