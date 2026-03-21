package db

import (
	"context"
	"database/sql"
)

type DB struct {
	Database Database
}

type Database interface {
	Connect() (*sql.DB, error)
	InitDB(ctx context.Context, db *sql.DB) error
}

func NewDB(db Database) *DB {
	return &DB{Database: db}
}
