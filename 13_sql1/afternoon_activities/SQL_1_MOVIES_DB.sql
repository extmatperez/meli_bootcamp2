/* 
PRIMERA PARTE

Abrir desde MySql Workbench el archivo movies_db.sql y ejecutar su contenido.

SEGUNDA PARTE
 */

-- 1) Mostrar todos los registros de la tabla de movies.
SELECT * FROM movies_db.movies LIMIT 0, 1000	21 row(s) returned	0.00047 sec / 0.000020 sec
-- 2) Mostrar el nombre, apellido y rating de todos los actores.
SELECT first_name, last_name, rating FROM movies_db.actors LIMIT 0, 1000	49 row(s) returned	0.00056 sec / 0.000025 sec
-- 3) Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español
SELECT title as título FROM movies_db.series LIMIT 0, 1000	6 row(s) returned	0.00039 sec / 0.000012 sec
-- 4) Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.
SELECT first_name, last_name, rating FROM movies_db.actors WHERE rating > 7.5 LIMIT 0, 1000	6 row(s) returned	0.00050 sec / 0.000010 sec
-- 5) Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.
SELECT title, rating, awards FROM movies_db.movies WHERE rating > 7.5 AND awards > 2 LIMIT 0, 1000	9 row(s) returned	0.00052 sec / 0.000014 sec
-- 6) Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
SELECT title, rating FROM movies_db.movies ORDER BY rating LIMIT 0, 1000	21 row(s) returned	0.00052 sec / 0.000021 sec
-- 7) Mostrar los títulos de las primeras tres películas en la base de datos.
SELECT title FROM movies_db.movies LIMIT 3	3 row(s) returned	0.00038 sec / 0.000011 sec
-- 8) Mostrar el top 5 de las películas con mayor rating.
SELECT title, rating FROM movies_db.movies ORDER BY rating desc LIMIT 5	5 row(s) returned	0.00061 sec / 0.000010 sec
-- 9) Mostrar las top 5 a 10 de las películas con mayor rating.
SELECT title, rating FROM movies_db.movies ORDER BY rating desc LIMIT 5 OFFSET 5	5 row(s) returned	0.00039 sec / 0.000010 sec
-- 10) Listar los primeros 10 actores (sería la página 1).
SELECT first_name, last_name FROM movies_db.actors LIMIT 10	10 row(s) returned	0.00056 sec / 0.000017 sec
-- 11) Luego usar offset para traer la página 3
SELECT first_name, last_name FROM movies_db.actors LIMIT 10 OFFSET 20	10 row(s) returned	0.00039 sec / 0.000011 sec
-- 12) Hacer lo mismo para la página 5
SELECT first_name, last_name FROM movies_db.actors LIMIT 10 OFFSET 40	9 row(s) returned	0.00047 sec / 0.000011 sec
-- 13) Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
SELECT title, rating FROM movies_db.movies WHERE title LIKE '%Toy Story%' LIMIT 0, 1000	2 row(s) returned	0.00058 sec / 0.000011 sec
-- 14) Mostrar a todos los actores cuyos nombres empiecen con Sam.
SELECT first_name, last_name FROM movies_db.actors WHERE first_name LIKE 'Sam%' LIMIT 0, 1000	2 row(s) returned	0.00063 sec / 0.000014 sec
-- 15) Mostrar el título de las películas que salieron entre el 2004 y 2008.
SELECT title, release_date FROM movies_db.movies WHERE release_date > '2003-12-31 00:00:00' AND release_date < '2009-01-01 00:00:00' LIMIT 0, 1000	7 row(s) returned	0.00048 sec / 0.000012 sec
SELECT title FROM movies_db.movies WHERE YEAR(release_date) BETWEEN '2004' AND '2008' LIMIT 0, 1000	7 row(s) returned	0.00066 sec / 0.000014 sec
-- 16) Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. 
--     Ordenar los resultados por rating.
SELECT title, rating FROM movies_db.movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN '1988' AND '2009' ORDER BY rating desc LIMIT 0, 1000	14 row(s) returned	0.00057 sec / 0.000012 sec
-- 17) Traer el top 3 a partir del registro 10 de la consulta anterior.
SELECT title, rating FROM movies_db.movies WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN '1988' AND '2009' ORDER BY rating desc LIMIT 3 OFFSET 9	3 row(s) returned	0.00054 sec / 0.000010 sec




