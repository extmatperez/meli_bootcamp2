#Ejercicio 2
SELECT * FROM movies;

#Ejercicio 3
SELECT first_name, last_name, rating FROM actors;

#Ejercicio 4
SELECT title AS titulo FROM series;

#Ejercicio 5
SELECT first_name as Nombre, last_name as Apellido FROM actors
WHERE rating > 7.5;

#Ejercicio 6
SELECT title as Titulo, rating as Rating, awards as Premios FROM movies
WHERE rating > 7.5
AND awards > 2;

#Ejercicio 7
SELECT title as Titulo, rating as Rating FROM movies
ORDER BY rating;

#Ejercicio 8
SELECT * FROM movies
LIMIT 3;

#Ejercicio 9
SELECT * FROM movies
ORDER BY rating DESC
LIMIT 5;

#Ejercicio 10
SELECT * FROM movies
ORDER BY rating DESC
LIMIT 5 OFFSET 5;

#Ejercicio 11
SELECT * FROM actors
LIMIT 10;

#Ejercicio 12
SELECT * FROM actors
LIMIT 10 OFFSET 20;

#Ejercicio 13
SELECT * FROM actors
LIMIT 10 OFFSET 40;

#Ejercicio 14
SELECT title as Titulo, rating as Rating FROM movies
WHERE title LIKE "%Toy Story%";

#Ejercicio 15
SELECT * FROM actors
WHERE first_name LIKE "Sam%";

#Ejercicio 16
SELECT * FROM movies
WHERE release_date BETWEEN "2004-01-01 00:00:00" AND "2009-01-01 00:00:00";

#Ejercicio 17 
SELECT title as Titulo FROM movies
WHERE rating > 3 
AND awards > 1 
AND release_date BETWEEN "1998-01-01 00:00:00" AND "2009-12-12 23:59:59";

#Ejercicio 18
SELECT * FROM movies
WHERE rating > 3 
AND awards > 1 
AND release_date BETWEEN "1998-01-01 00:00:00" AND "2009-12-12 23:59:59"
ORDER BY rating DESC
LIMIT 3;