package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBPool struct {
	Pool *pgxpool.Pool
}

func NewConn() (*DBPool, error) {
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		return &DBPool{}, fmt.Errorf("The db connection string in env variables not found")
	}

	dbpool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return &DBPool{}, fmt.Errorf("Error while connectng to db: %w", err)
	}

	if err := dbpool.Ping(context.Background()); err != nil {
		return &DBPool{}, fmt.Errorf("Failed to ping the database: %w", err)
	}

	return &DBPool{Pool: dbpool}, nil
}

func (db *DBPool) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
