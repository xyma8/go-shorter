package db

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

type Sqlite struct {
}

func NewSqlite() *Sqlite {
	return &Sqlite{}
}

func (s *Sqlite) Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./go-shorter.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func (s *Sqlite) InitDB(ctx context.Context, db *sql.DB) error {
	const schema = `
	PRAGMA foreign_keys = ON;

	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		original_url TEXT NOT NULL,
		short_url TEXT UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.ExecContext(ctx, schema)
	return err
}
