package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() {
	connStr := "postgres://postgres:mamayogho@localhost:5432/patient_db"

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Unable to create connection pool: ", err)
	}

	Pool = pool
	log.Println("Connected to patient_db")
}
