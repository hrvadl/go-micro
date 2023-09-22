package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	*sql.DB
}

func MustConnect(dsn string) *Database {
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	return &Database{db}
}
