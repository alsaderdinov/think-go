package db

import "context"

type Movie struct {
	ID          int     `db:"id"`
	StudioID    int     `db:"studio_id"`
	Title       string  `db:"title"`
	ReleaseYear int     `db:"release_year"`
	BoxOffice   float64 `db:"box_office"`
	Rating      string  `db:"rating"`
}

type Repository interface {
	AddMovies(ctx context.Context, movies []Movie) error
	DeleteMovie(ctx context.Context, id int) error
	UpdateMovie(ctx context.Context, movie Movie) error
	Movies(ctx context.Context) ([]Movie, error)
}
