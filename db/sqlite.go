package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

func Connect() (*DB, error) {
	var user string = os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	sslMode := os.Getenv("POSTGRES_SSLMODE")
	if sslMode == "" {
		sslMode = "disable"
	}

	connectStr := []string{"postgresql://", user, ":", password, "@", host, ":", port, "/", dbName, "?sslmode=", sslMode}
	db, err := sql.Open("postgres", strings.Join(connectStr, ""))
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	fmt.Println("Успешно подключено к PostgreSQL!")

	return &DB{db}, nil
}

func Init(ctx context.Context, db *sql.DB) error {
	const schema = `
	CREATE TABLE IF NOT EXISTS urls (
		id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		original_url TEXT NOT NULL,
		short_url TEXT UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.ExecContext(ctx, schema)
	return err
}
