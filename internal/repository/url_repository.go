package repository

import (
	_ "database/sql"

	"github.com/xyma8/go-shorter/db"
)

type UrlRepository struct {
	*db.DB
}

func NewUrlRepository(db *db.DB) *UrlRepository {
	return &UrlRepository{db}
}

func InsertUrl() {

}
