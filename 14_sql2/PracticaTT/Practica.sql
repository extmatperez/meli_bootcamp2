#Explicar el concepto de normalización y para que se utiliza.
#Es un proceso de estandarización y validación de datos en el cual se eliminan redundacias e inconsistencias, con la finalidad de que los datos sean mas simples de consultar y más eficiente para el que la gestiona.

#Agregar una película a la tabla movies.
SELECT * FROM movies;

INSERT INTO movies(title, rating, awards, release_date, length, genre_id)
values ("Toy Story 3", 8.4, 6, "2010-06-15 00:00:00", 200, 7);

#Agregar un género a la tabla genres.
SELECT * FROM genres;

INSERT INTO genres(created_at, name, ranking, active)
values ("2021-12-16 16:35:00","Policial", 14, 1);

#Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies
SET genre_id = 15
WHERE id > 22;

#Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors
SET favorite_movie_id = 22
WHERE first_name = "Jon";

#Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE temp_movies
(SELECT * FROM movies);

#Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM temp_movies
WHERE awards<5;

SELECT * FROM temp_movies;

#Obtener la lista de todos los géneros que tengan al menos una película.
SELECT DISTINCT genres.name
FROM movies INNER JOIN genres
ON movies.genre_id = genres.id;

#Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.
SELECT actors.first_name, actors.last_name
FROM actors INNER JOIN movies
ON actors.favorite_movie_id = movies.id
WHERE movies.awards > 3;

SELECT * FROM actors;

#Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

#¿Qué son los índices? ¿Para qué sirven?
#Los índices son un mecanismo que se utiliza para optimizar las consultas en SQL. Otorgan una ruta más rápida a los registros evitando barrer todos los valores de un campo.

#Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX name_idx
ON movies (title);

#Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;