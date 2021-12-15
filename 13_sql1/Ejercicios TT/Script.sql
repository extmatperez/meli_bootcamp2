SELECT * FROM movies;


SELECT nombre, apellido, rating FROM actores;

SELECT nombre, apellido FROM actores WHERE rating > 7.5;

SELECT name AS nombre FROM series as series


SELECT titulos, rating, premios from movies where premios > 2 AND rating > 7.5

SELECT titulo, rating from peliculas ORDER BY rating ASC 

SELECT titulo from peliculas limit 3

select titulo, rating FROM peliculas ORDER BY rating desc LIMIT 5

select titulo, rating FROM peliculas ORDER BY rating desc LIMIT 5 OFFSET 5

SELECT * FROM  actores LIMIT 10

SELECT * FROM  actores LIMIT 10 OFFSET 30

SELECT * FROM  actores LIMIT 10 OFFSET 50

select titulo, rating FROM peliculas WHERE titulo LIKE '%Toy Story%'

select titulo, rating FROM peliculas WHERE titulo LIKE 'Sam%'

select titulo FROM peliculas WHERE año BETWEEN 2004 and 2008

select titulo FROM peliculas WHERE rating > 3 and premios > 1 and año BETWEEN 1988 and 2009 order by rating asc

select titulo FROM peliculas WHERE rating > 3 and premios > 1 and año BETWEEN 1988 and 2009 order by rating desc LIMIT 3






