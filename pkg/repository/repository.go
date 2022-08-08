package repository

import (
	"github.com/jmoiron/sqlx"

	"astro"
)

type Picture interface {
	Insert(p astro.Picture) (int64, error)
	GetByDate(date string) ([]astro.Picture, error)
	GetByDateRange(start, end string) ([]astro.Picture, error)
}

type Repository struct {
	Picture
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Picture: NewPostgres(db),
	}
}
