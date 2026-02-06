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
	UpdateShortUrl(ctx context.Context, id uint, url string) error
	GetOrigUrl(ctx context.Context, shortUrl string) (string, error)
}

func NewUrlService(repo UrlRepository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) ShortenUrl(ctx context.Context, url *models.UrlModel) (string, error) {
	id, err := s.repo.CreateUrl(ctx, url)
	if err != nil {
		return "", err
	}

	obfId, err := helpers.EncodeBiject(id)
	if err != nil {
		return "", err
	}
	//fmt.Println(obfId)

	shortUrl, err := helpers.EncodeURLBase62(uint(obfId), 5)
	if err != nil {
		return "", err
	}

	err = s.repo.UpdateShortUrl(ctx, id, shortUrl)
	if err != nil {
		return "", err
	}

	return shortUrl, nil
}

func (s *UrlService) GetOrigUrl(ctx context.Context, shortUrl string) (string, error) {
	url, err := s.repo.GetOrigUrl(ctx, shortUrl)
	if err != nil {
		return "", err
	}

	return url, nil
}
