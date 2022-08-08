package service

import (
	"astro/pkg/repository"

	"astro"
)

type Picture interface {
	Insert(p astro.Picture) (int64, error)
	GetByDate(date string) ([]astro.Picture, error)
	GetByDateRange(start, end string) ([]astro.Picture, error)
}

type Service struct {
	Picture
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Picture: NewAstroService(repos.Picture),
	}
}
