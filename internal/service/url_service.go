package service

import (
	"context"

	"github.com/xyma8/go-shorter/internal/models"
)

type UrlService struct {
	repo UrlRepository
}

type UrlRepository interface {
	CreateUrl(ctx context.Context, url *models.UrlModel) (*models.UrlModel, error)
}

func NewUrlService(repo UrlRepository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) ShortUrl(ctx context.Context, url *models.UrlModel) (*models.UrlModel, error) {
	return s.repo.CreateUrl(ctx, url)
}
