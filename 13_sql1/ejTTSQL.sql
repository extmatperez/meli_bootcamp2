USE movies_db;
SELECT * FROM movies;

SELECT nombre,apellido,rating FROM actors;

SELECT title AS Titulo  FROM series;

SELECT first_name, last_name FROM actors WHERE rating > 7.5;

SELECT title,rating,awards FROM movies WHERE rating > 7.5 AND awards > 2;

SELECT title,rating FROM movies order by rating ASC;

SELECT title FROM movies LIMIT 3;

SELECT title,rating FROM movies order by rating DESC LIMIT 5;

SELECT title,rating FROM movies order by rating DESC LIMIT 5 offset 5;

SELECT * FROM actors LIMIT 10;

SELECT * FROM actors LIMIT 10 offset 20;

SELECT * FROM actors LIMIT 10 offset 40;

SELECT title,rating FROM movies WHERE title LIKE '%Toy Story%';

SELECT * FROM actors WHERE first_name LIKE 'Sam%';

SELECT title FROM movies WHERE year(release_date) BETWEEN 2004 AND 2008;

SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND year(release_date) BETWEEN 1988 AND 2009 order by rating;

SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND year(release_date) BETWEEN 1988 AND 2009 order by rating LIMIT 3 OFFSET 10;
