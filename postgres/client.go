package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jmoiron/sqlx"
)

func NewPGXConnection() (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return db, nil
}

func NewSQLXConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return db, nil
}

func NewSQLConnection() (*sql.DB, error) {
	db, err := sql.Open("", "")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return db, nil
}
