package service

import (
	"context"

	"github.com/xyma8/go-shorter/internal/models"
)

type UrlService struct {
	repo UrlRepository
}

type UrlRepository interface {
	CreateUrl(ctx context.Context, url *models.UrlModel) (uint, error)
	UpdateUrl(ctx context.Context, id uint, url *models.UrlModel) error
}

func NewUrlService(repo UrlRepository) *UrlService {
	return &UrlService{repo: repo}
}

func (s *UrlService) ShortUrl(ctx context.Context, url *models.UrlModel) (string, error) {
	id, err := s.repo.CreateUrl(ctx, url)
	if err != nil {
		return "", err
	}

	obfId, err := encodeXOR(id)
	if err != nil {
		return "", err
	}

	shortUrl, err := encodeBase62(obfId)
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

func encodeBase62(id uint) (string, error) {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	binary_result := []byte{0, 0, 0, 0, 0}

	if id >= uintPow(62, 5) {
		//return "", err
	}

	for i := 1; id > 0; i++ {
		mod := id % 62
		id = id / 62
		binary_result[len(binary_result)-i] = alphabet[mod]
	}
	return string(binary_result), nil
}

func encodeXOR(id uint) (uint, error) {
	const MASK uint = 0x2A5B8D3F
	id = id ^ MASK
	return id, nil
}

func uintPow(base, exponent uint) uint {
	result := uint(1)
	for i := uint(0); i < exponent; i++ {
		result *= base
	}

	return result
}
