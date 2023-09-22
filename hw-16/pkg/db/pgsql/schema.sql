DROP TABLE IF EXISTS movies_directors;
DROP TABLE IF EXISTS movies_actors;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS directors;
DROP TABLE IF EXISTS studios;
DROP TYPE IF EXISTS rating;

-- Create the studios table
CREATE TABLE studios (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

-- Create the directors table
CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    birth_date DATE NOT NULL
);

-- Create the actors table
CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    birth_date DATE NOT NULL
);

-- Define the rating ENUM
CREATE TYPE rating AS ENUM ('PG-10', 'PG-13', 'PG-18');

-- Create the movies table
CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    release_year INT CHECK (release_year >= 1800),
    box_office DECIMAL(12, 2),
    rating rating,
    studio_id INT DEFAULT 0 REFERENCES studios(id),
    UNIQUE(title, release_year)
);

-- Create the relation between movies and actors
CREATE TABLE movies_actors (
    movie_id INT REFERENCES movies(id),
    actor_id INT REFERENCES actors(id),
    PRIMARY KEY(movie_id, actor_id)
);

-- Create the relation between movies and directors
CREATE TABLE movies_directors (
    movie_id INT REFERENCES movies(id),
    director_id INT REFERENCES directors(id),
    PRIMARY KEY(movie_id, director_id)
);
