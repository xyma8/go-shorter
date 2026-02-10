package service

import (
	"context"

	"github.com/xyma8/go-shorter/internal/helpers"
	"github.com/xyma8/go-shorter/internal/models"
)

type UrlService struct {
	repo UrlRepository
}

type UrlRepository interface {
	CreateUrl(ctx context.Context, url *models.CreatingUrl) (uint, error)
	UpdateShortUrl(ctx context.Context, id uint, url string) error
	GetOrigUrl(ctx context.Context, shortUrl string) (string, error)
}

func NewUrlService(repo UrlRepository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) ShortenUrl(ctx context.Context, url *models.CreatingUrl, permuteKey string) (*models.Url, error) {
	id, err := s.repo.CreateUrl(ctx, url)
	if err != nil {
		return nil, err
	}

	obfId, err := helpers.PermuteRange(uint64(id), []byte(permuteKey))
	if err != nil {
		return nil, err
	}

	shortUrl, err := helpers.EncodeURLBase62(uint(obfId), 5)
	if err != nil {
		return nil, err
	}

	err = s.repo.UpdateShortUrl(ctx, id, shortUrl)
	if err != nil {
		return nil, err
	}

	resUrl := models.Url{
		Original_url: url.Original_url,
		Short_url:    shortUrl,
	}
	return &resUrl, nil
}

func (s *UrlService) GetOrigUrl(ctx context.Context, shortUrl string) (*models.OrigUrl, error) {
	url, err := s.repo.GetOrigUrl(ctx, shortUrl)
	if err != nil {
		return nil, err
	}

	origUrl := models.OrigUrl{
		Original_url: url,
	}

	return &origUrl, nil
}
