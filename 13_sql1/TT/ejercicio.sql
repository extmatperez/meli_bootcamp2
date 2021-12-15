# Ejercicio 2
SELECT * FROM movies;

# Ejercicio 3
SELECT first_name, last_name, rating FROM actors;

# Ejercicio 4
SELECT title as titulo FROM series serie;

# Ejercicio 5
SELECT first_name, last_name FROM actors WHERE rating > 7.5;

# Ejercicio 6
SELECT title, rating, awards FROM movies WHERE rating > 7.5 AND awards > 2;

# Ejercicio 7
SELECT title, rating FROM movies WHERE rating > 7.5 ORDER BY rating ASC;

# Ejercicio 8
SELECT title FROM movies LIMIT 3;

# Ejercicio 9
SELECT * FROM movies ORDER BY rating DESC LIMIT 5;

# Ejercicio 10
SELECT * FROM movies ORDER BY rating DESC LIMIT 5 OFFSET 5;

# Ejercicio 11
SELECT * FROM actors LIMIT 10 OFFSET 0;

# Ejercicio 12
SELECT * FROM actors LIMIT 10 OFFSET 20;

# Ejercicio 13
SELECT * FROM actors LIMIT 10 OFFSET 40;

# Ejercicio 14
SELECT title, rating FROM movies WHERE title = "Toy Story";

# Ejercicio 15
SELECT first_name FROM actors WHERE first_name LIKE "Sam%";

# Ejercicio 16
SELECT title FROM movies WHERE YEAR(release_date) BETWEEN 2004 AND 2008;

# Ejercicio 17
SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN 1988 AND 2009 ORDER BY rating DESC;

# Ejercicio 18
SELECT title FROM movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN 1988 AND 2009 ORDER BY rating DESC LIMIT 3 OFFSET 10;


