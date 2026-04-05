package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return db, nil
}
