package repository

import (
	"context"
	"database/sql"

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

func (r *urlRepository) CreateUrl(ctx context.Context, url *models.CreatingUrl) (uint, error) {
	query := "INSERT INTO urls (original_url) VALUES ($1) RETURNING id"
	var insertID int
	err := r.db.QueryRow(query, url.Original_url).Scan(&insertID)
	if err != nil {
		return 0, err
	}

	//insertID, err := result.LastInsertId()
	//if err != nil {
	//	return 0, err
	//}

	return uint(insertID), nil
}

func (r *urlRepository) UpdateShortUrl(ctx context.Context, id uint, url string) error {
	query := "UPDATE urls SET short_url = $1 WHERE id = $2"
	_, err := r.db.Exec(query, url, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *urlRepository) GetOrigUrl(ctx context.Context, shortUrl string) (string, error) {
	query := "SELECT original_url FROM urls WHERE short_url = $1"
	var res string
	err := r.db.QueryRow(query, shortUrl).Scan(&res)
	if err != nil {
		return "", err
	}

	return res, nil
}
