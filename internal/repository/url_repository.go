package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/xyma8/go-shorter/internal/models"
)

type urlRepository struct {
	db *sql.DB
}

//type UrlRepository interface {
//Create(ctx context.Context, url *models.UrlModel) (*models.UrlModel, error)
//}

func NewUrlRepository(db *sql.DB) *urlRepository {
	return &urlRepository{db}
}

func (r *urlRepository) CreateUrl(ctx context.Context, url *models.UrlModel) (*models.UrlModel, error) {
	fmt.Print("repository")
	fmt.Print(url.Original_url)
	return nil, nil
}
