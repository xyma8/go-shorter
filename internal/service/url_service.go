package service

import (
	"github.com/xyma8/go-shorter/internal/repository"
)

type UrlService struct {
}

func NewUrlService(repo *repository.UrlRepository) *UrlService {
	return &UrlService{}
}
