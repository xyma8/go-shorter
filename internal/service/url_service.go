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
	CreateUrl(ctx context.Context, url *models.UrlModel) (uint, error)
	UpdateUrl(ctx context.Context, id uint, url *models.UrlModel) error
	GetOrigUrl(ctx context.Context, shortUrl string) (string, error)
}

func NewUrlService(repo UrlRepository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) ShortUrl(ctx context.Context, url *models.UrlModel) (string, error) {
	id, err := s.repo.CreateUrl(ctx, url)
	if err != nil {
		return "", err
	}

	obfId, err := helpers.EncodeFeistel(id)
	if err != nil {
		return "", err
	}

	shortUrl, err := helpers.EncodeBase62(obfId)
	if err != nil {
		return "", err
	}

	url.Short_url = shortUrl
	err = s.repo.UpdateUrl(ctx, id, url)
	if err != nil {
		return "", err
	}

	return url.Short_url, nil
}

func (s *UrlService) GetOrigUrl(ctx context.Context, shortUrl string) (string, error) {
	url, err := s.repo.GetOrigUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}

	return url, nil
}

func encodeBiject(n uint) (uint, error) {
	const M uint = 123412340
	const a uint = 198134563
	const b uint = 10

	return (a*n + b) % M, nil
}
