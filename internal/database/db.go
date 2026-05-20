package database

import (
	"context"
	"golang-CEX/internal/database/sqlc"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	Pool    *pgxpool.Pool
	Queries sqlc.Queries
}

func New() (*Store, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := pgxpool.New(ctx, "postgres://admin:password@localhost:5432/mydb?sslmode=disable")
	if err != nil {
		log.Fatal("The error in db pool connection is", err)
		return nil, err
	}
	if err = conn.Ping(ctx); err != nil {
		log.Fatal("The error in pinging the connection is :", err)
		return nil, err
	}

	return &Store{
		Pool:    conn,
		Queries: *sqlc.New(conn),
	}, nil

}
