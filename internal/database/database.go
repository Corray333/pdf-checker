package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	Db *sql.DB
}

func ConnectToDatabase() *DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_CONNECTIOIN_STR"))
	if err != nil {
		log.Fatal(err)
	}

	var result DB
	result.Db = db

	return &result
}
