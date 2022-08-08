package repository

import (
	"astro"
	"time"

	"github.com/jmoiron/sqlx"
)

type Actions struct {
	db *sqlx.DB
}

func NewPostgres(db *sqlx.DB) *Actions {
	return &Actions{db: db}
}

func (r *Actions) Insert(p astro.Picture) (int64, error) {

	query := `INSERT INTO pictures 
	(copyright, explanation, hd_url, media_type, service_version, title, url, binary_pic, pic_of_the_day, stored)
	values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	row, err := r.db.Exec(query, p.Copyright, p.Explanation, p.Hdurl, p.MediaType, p.ServiceVersion, p.Title, p.Url, p.BinaryPic, p.PicOfTheDay, time.Now().Format("2006-01-02"))
	if err != nil {
		return 0, err
	}

	return row.RowsAffected()
}

func (r *Actions) GetByDate(date string) ([]astro.Picture, error) {
	var picture []astro.Picture
	if err := r.db.Select(&picture,
		`SELECT copyright, explanation, hd_url, media_type, service_version, title, url, pic_of_the_day, stored
		 FROM pictures WHERE pic_of_the_day = $1`, date); err != nil {
		return nil, err
	}

	return picture, nil
}

func (r *Actions) GetByDateRange(start, end string) ([]astro.Picture, error) {
	var pictures []astro.Picture
	if err := r.db.Select(&pictures,
		`SELECT copyright, explanation, hd_url, media_type, service_version, title, url, binary_pic, pic_of_the_day, stored
		 FROM pictures WHERE pic_of_the_day > $1 AND pic_of_the_day < $2`, start, end); err != nil {

		return nil, err
	}
	return pictures, nil
}
