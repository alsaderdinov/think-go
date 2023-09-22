package pgsql

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"
	"think-go/hw-16/pkg/db"
)

var testDB *DB

func TestMain(m *testing.M) {
	initDB()

	code := m.Run()
	os.Exit(code)
}

func initDB() {
	var err error

	testDB, err = New("host=localhost database=kinopoisk_test")
	if err != nil {
		log.Fatal(err)
	}

	execSQL("schema.sql")
	execSQL("data.sql")
}

func execSQL(f string) {
	data, err := os.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	_, err = testDB.pool.Exec(context.Background(), string(data))
	if err != nil {
		log.Fatal(err)
	}
}

func TestDB_Movies(t *testing.T) {
	tests := []struct {
		name     string
		studioID int
		want     []db.Movie
	}{
		{
			name:     "Test with studio_id",
			studioID: 1,
			want: []db.Movie{
				{
					ID:          1,
					Title:       "Fight Club",
					ReleaseYear: 1999,
					BoxOffice:   100.85,
					Rating:      "PG-18",
					StudioID:    1,
				},
			},
		},
		{
			name:     "Test without studio_id",
			studioID: 0,
			want: []db.Movie{
				{
					ID:          1,
					Title:       "Fight Club",
					ReleaseYear: 1999,
					BoxOffice:   100.85,
					Rating:      "PG-18",
					StudioID:    1,
				},
				{
					ID:          2,
					Title:       "Whiplash",
					ReleaseYear: 2013,
					BoxOffice:   48.98,
					Rating:      "PG-18",
					StudioID:    2,
				},
				{
					ID:          3,
					Title:       "The Dark Knight",
					ReleaseYear: 2008,
					BoxOffice:   1003.04,
					Rating:      "PG-13",
					StudioID:    3,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDB.Movies(context.Background(), tt.studioID)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestDB_AddMovies(t *testing.T) {
	movies := []db.Movie{
		{
			Title:       "The Matrix",
			ReleaseYear: 1999,
			BoxOffice:   463.51,
			Rating:      "PG-18",
			StudioID:    3,
		},
		{
			Title:       "The Matrix Reloaded",
			ReleaseYear: 2003,
			BoxOffice:   742.12,
			Rating:      "PG-18",
			StudioID:    3,
		},
	}

	if err := testDB.AddMovies(context.Background(), movies); err != nil {
		t.Errorf("got error %v when trying to add movies", err)
	}

	got, err := testDB.Movies(context.Background(), 3)
	if err != nil {
		t.Errorf("got error %v when trying to retrieve movies", err)
	}

	want := []db.Movie{
		{
			ID:          3,
			Title:       "The Dark Knight",
			ReleaseYear: 2008,
			BoxOffice:   1003.04,
			Rating:      "PG-13",
			StudioID:    3,
		},
		{
			ID:          4,
			Title:       "The Matrix",
			ReleaseYear: 1999,
			BoxOffice:   463.51,
			Rating:      "PG-18",
			StudioID:    3,
		},
		{
			ID:          5,
			Title:       "The Matrix Reloaded",
			ReleaseYear: 2003,
			BoxOffice:   742.12,
			Rating:      "PG-18",
			StudioID:    3,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDB_UpdateMovie(t *testing.T) {
	updatedMovie := db.Movie{
		ID:          1,
		Title:       "Fight Club",
		ReleaseYear: 1999,
		BoxOffice:   150.00,
		Rating:      "PG-13",
		StudioID:    1,
	}

	if err := testDB.UpdateMovie(context.Background(), updatedMovie); err != nil {
		t.Errorf("got error %v on update movie", err)
	}

	got, err := testDB.Movies(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got[0], updatedMovie) {
		t.Errorf("got %v want %v", got[0], updatedMovie)
	}
}

func TestDB_DeleteMovie(t *testing.T) {
	want := []db.Movie{}

	if err := testDB.DeleteMovie(context.Background(), 1); err != nil {
		t.Errorf("got %v on delete movie", err)
	}

	got, err := testDB.Movies(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
