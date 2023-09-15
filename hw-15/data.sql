-- Insert studios
INSERT INTO studios (name)
VALUES
    ('Regency Enterprises'),
    ('Bold Films'),
    ('Warner Bros. Pictures');

-- Insert directors
INSERT INTO directors (name, birth_date)
VALUES
    ('David Fincher', '1962-08-28'),
    ('Damien Chazelle', '1985-01-19'),
    ('Christopher Nolan', '1970-07-30');

-- Insert actors
INSERT INTO actors (name, birth_date)
VALUES
    ('Edward Norton', '1969-08-18'),
    ('Brad Pitt', '1963-12-18'),
    ('Helena Bonham Carter', '1966-05-26'),
    ('Miles Teller', '1987-02-20'),
    ('J.K. Simmons', '1955-01-09'),
    ('Christian Bale', '1974-01-30'),
    ('Heath Ledger', '1979-04-04'),
    ('Aaron Eckhart', '1968-03-12');

-- Insert movies
INSERT INTO movies (title, release_year, box_office, rating, studio_id)
VALUES
    ('Fight Club', 1999, 100.85, 'PG-18', 1),
    ('Whiplash', 2013, 48.98, 'PG-18', 2),
    ('The Dark Knight', 2008, 1003.04, 'PG-13', 3);

-- Insert movie_actors
INSERT INTO movies_actors (movie_id, actor_id)
VALUES
    (1, 1),
    (1, 2),
    (1, 3),
    (2, 4),
    (2, 5),
    (3, 6),
    (3, 7),
    (3, 8);

-- Insert movie_directors
INSERT INTO movies_directors (movie_id, director_id)
VALUES
    (1, 1),
    (2, 2),
    (3, 3);
