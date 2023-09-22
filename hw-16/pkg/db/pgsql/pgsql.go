package pgsql

import (
	"context"
	"think-go/hw-16/pkg/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB represents the Postgres database connection.
type DB struct {
	pool *pgxpool.Pool
}

// New creates a new DB instance with the given connection string.
func New(conn string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		return nil, err
	}

	return &DB{pool: pool}, nil
}

// Movies retrieves a list of movies from the database based on the provided studioID.
func (d *DB) Movies(ctx context.Context, studioID int) ([]db.Movie, error) {
	rows, err := d.pool.Query(ctx,
		`SELECT id, studio_id, title, release_year, box_office, rating
		FROM movies
		WHERE $1 = 0 OR studio_id = $1
	`,
		studioID,
	)
	if err != nil {
		return nil, err
	}

	movies, err := pgx.CollectRows(rows, pgx.RowToStructByName[db.Movie])
	if err != nil {
		return nil, err
	}

	return movies, nil
}

// AddMovies adds a list of movies to the database in a transactional manner.
func (d *DB) AddMovies(ctx context.Context, movies []db.Movie) error {
	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	batch := &pgx.Batch{}

	for _, m := range movies {
		batch.Queue(`
			INSERT INTO movies(studio_id, title, release_year, box_office, rating)
			VALUES($1, $2, $3, $4, $5)
		`, m.StudioID, m.Title, m.ReleaseYear, m.BoxOffice, m.Rating)
	}

	res := tx.SendBatch(ctx, batch)

	err = res.Close()
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// UpdateMovie updates a movie in the database.
func (d *DB) UpdateMovie(ctx context.Context, movie db.Movie) error {
	_, err := d.pool.Exec(ctx, `
		UPDATE movies
		SET title = $1, release_year = $2, box_office = $3, rating = $4, studio_id = $5
		WHERE id = $6
	`, movie.Title, movie.ReleaseYear, movie.BoxOffice, movie.Rating, movie.StudioID, movie.ID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteMovie deletes a movie from the database based on the provided ID.
func (d *DB) DeleteMovie(ctx context.Context, id int) error {
	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	queries := []string{
		"DELETE FROM movies_actors WHERE movie_id = $1",
		"DELETE FROM movies_directors WHERE movie_id = $1",
		"DELETE FROM movies WHERE id = $1",
	}

	for _, q := range queries {
		_, err = tx.Exec(ctx, q, id)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
