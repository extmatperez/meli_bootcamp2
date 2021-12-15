USE movies_db;
SELECT * FROM movies;
SELECT first_name, last_name, rating FROM actors;
SELECT title AS titulo FROM series AS Serie;
SELECT first_name, last_name FROM actors
WHERE rating > 7.5;
SELECT title, rating FROM movies ORDER BY rating ASC;
SELECT title FROM movies LIMIT 3;
SELECT * FROM movies ORDER BY rating DESC LIMIT 5;
SELECT * FROM movies ORDER BY rating DESC LIMIT 5 OFFSET 6;
SELECT * FROM actors LIMIT 10;
SELECT * FROM actors LIMIT 10 OFFSET 20; 
SELECT * FROM actors LIMIT 10 OFFSET 40;
SELECT title, rating FROM movies WHERE title LIKE "Toy Story%";
SELECT * FROM actors WHERE first_name LIKE "Sam%";
SELECT title FROM movies WHERE YEAR(release_date) BETWEEN '2004' AND '2008';
SELECT * FROM movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN '1988' AND '2009';
SELECT * FROM movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN '1988' AND '2009' LIMIT 3 OFFSET 9;