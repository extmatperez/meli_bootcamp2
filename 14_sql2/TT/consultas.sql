# 1. Explicar el concepto de normalización y para que se utiliza.
/*
    La normalizacion es un proceso que tiene como objetivo eliminar la redundancia e inconsistencias de forma tal que la integridad de la información este protegida,
    a su vez, se favorece su interpretacion, por ultima que sea mas facil de consultar y de gestionar.
*/

# 2. Agregar una película a la tabla movies.
INSERT INTO movies (title, rating, awards, release_date) VALUES ('Cars', 8.5, 1, '2010-01-10 00:00:00');

# 3. Agregar un género a la tabla genres.
INSERT INTO genres (name, ranking, active) VALUES ('Autitos', 20, 1);

# 4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies
SET genre_id = 13
WHERE movies.id = 22;

# 5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors
SET favorite_movie_id = 22
WHERE id = 1;

# 6. Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_copy
SELECT *
FROM movies;

# 7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM movies_copy
WHERE awards < 5;
#TODO: no me funca

# 8. Obtener la lista de todos los géneros que tengan al menos una película.
SELECT genres.*
FROM genres
JOIN movies ON movies.genre_id = genres.id;

# 9. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
SELECT actors.*
FROM actors
JOIN movies ON movies.id = actors.favorite_movie_id
WHERE movies.awards > 3;

# 10. Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
#TODO: me tira error

# 11. ¿Qué son los índices? ¿Para qué sirven?
/*
    Son mecanismos que se utilizan para optimizar consultas y permiten mejorar el acceso a los datos.
*/

# 12. Crear un índice sobre el nombre en la tabla movies.
CREATE UNIQUE INDEX movies_name_test ON movies(title);

# 13. Chequee que el índice fue creado correctamente.
INSERT INTO movies (title, rating, awards, release_date) VALUES ('Cars', 8.5, 1, '2010-01-10 00:00:00');
# Debe tirar error debido a que el titulo de la pelicula debe ser unico.
