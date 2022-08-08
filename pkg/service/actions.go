package service

import (
	"astro"

	"astro/pkg/repository"
)

type AstroService struct {
	repo repository.Picture
}

func NewAstroService(repo repository.Picture) *AstroService {
	return &AstroService{repo: repo}
}

func (s *AstroService) Insert(pic astro.Picture) (int64, error) {
	return s.repo.Insert(pic)
}

func (s *AstroService) GetByDate(date string) ([]astro.Picture, error) {
	return s.repo.GetByDate(date)
}

func (s *AstroService) GetByDateRange(start, end string) ([]astro.Picture, error) {
	return s.repo.GetByDateRange(start, end)
}
