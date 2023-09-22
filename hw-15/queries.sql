-- Selection of films with the name of the studio;
SELECT
    m.id AS movie_id,
    m.title AS movie_title,
    s.name AS studio_name
FROM
    movies m
JOIN
    studios s ON m.studio_id = s.id;

-- Selection of films for a certain actor;
SELECT
    m.title AS movie_title
FROM
    movies m
JOIN
    movies_actors ma ON m.id = ma.movie_id
JOIN
    actors a ON a.id = ma.actor_id
WHERE
    a.name = 'Edward Norton';

-- Count of films for a certain director;
SELECT
    COUNT(m.id) AS movies_count
FROM
    movies m
JOIN
    movies_directors md ON m.id = md.movie_id
JOIN
    directors d ON d.id = md.director_id
WHERE
    d.name = 'David Fincher';

-- Selection of films for several directors from the list;
SELECT
    m.title
FROM
    movies m
JOIN
    movies_directors md ON m.id = md.movie_id
JOIN
    directors d ON md.director_id = d.id
WHERE
    d.name IN ('Damien Chazelle', 'Christopher Nolan');

-- Count of films for a certain actor;
SELECT
    COUNT(m.id) AS movies_count
FROM
    movies m
JOIN
    movies_actors ma ON m.id = ma.movie_id
JOIN
    actors a ON a.id = ma.actor_id
WHERE
    a.name = 'J.K. Simmons';

-- Sampling of actors and directors who have appeared in more than 2 films;
-- Actors
SELECT
    a.name
FROM
    actors a
JOIN
    movies_actors ma ON a.id = ma.actor_id
GROUP BY
    a.name
HAVING
    COUNT(ma.movie_id) > 2;

-- Directors
SELECT
    d.name
FROM
    directors d
JOIN
    movies_directors md ON d.id = md.director_id
GROUP BY
    d.name
HAVING
    COUNT(md.movie_id) > 2;

-- Counting the number of films that have grossed over 1,000;
SELECT
    COUNT(m.id)
FROM
    movies m
WHERE
    m.box_office > 1000;

-- Counting the number of distinct movies that have grossed over 1,000;
SELECT
    COUNT(DISTINCT md.movie_id)
FROM
    movies m
JOIN
    movies_directors md ON m.id = md.movie_id
WHERE
    m.box_office > 1000;

-- Selection of distinct last names of actors;
SELECT DISTINCT
    SPLIT_PART(name, ' ', -1) AS last_name
FROM
    actors;

-- Counting the number of films that have duplicate titles;
SELECT
    m.title,
    COUNT(*) AS duplicate_count
FROM
    movies m
GROUP BY
    m.title
HAVING
    COUNT(*) > 1;
