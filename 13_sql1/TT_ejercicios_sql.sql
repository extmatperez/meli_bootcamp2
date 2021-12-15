SELECT * FROM movies_db.movies;

SELECT first_name, last_name, rating FROM actors;

SELECT title AS Titulo FROM series;

SELECT * FROM actors
WHERE rating > 7.5
ORDER BY rating;

SELECT title, rating, awards FROM movies
WHERE rating > 7.5 AND awards > 2;

SELECT title, rating FROM movies
ORDER BY rating;

SELECT title, id FROM movies
LIMIT 3;

SELECT title, rating FROM movies
ORDER BY rating DESC
LIMIT 5;

SELECT title, rating FROM movies
ORDER BY rating DESC
LIMIT 5 OFFSET 5;

SELECT first_name, last_name FROM actors
LIMIT 10;

SELECT first_name, last_name FROM actors
LIMIT 10 OFFSET 20;

SELECT first_name, last_name FROM actors
LIMIT 10 OFFSET 40;

SELECT title FROM movies
WHERE title = 'Toy Story';

SELECT title FROM movies
WHERE title LIKE 'Toy Story%';

SELECT first_name, last_name FROM actors
WHERE first_name LIKE 'Sam%';

SELECT title, release_date FROM movies
WHERE release_date BETWEEN '2004-01-01 00:00:00' AND '2008-12-31 23:59:59';

SELECT title, rating, awards, release_date FROM movies
WHERE release_date BETWEEN '1988-01-01 00:00:00' AND '2009-12-31 23:59:59'
AND rating > 3 AND awards > 1
ORDER BY rating;

SELECT title, rating, awards, release_date FROM movies
WHERE release_date BETWEEN '1988-01-01 00:00:00' AND '2009-12-31 23:59:59'
AND rating > 3 AND awards > 1
ORDER BY rating
LIMIT 3 OFFSET 10;
