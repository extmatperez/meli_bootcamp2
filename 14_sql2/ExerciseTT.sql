SELECT * FROM movies_db.movies;
SELECT * FROM movies_db.genres;
SELECT * FROM movies_db.actors;
SELECT * FROM movies_db.copy_movies;


#Agregar una película a la tabla movies.
INSERT INTO movies_db.movies (title, rating, release_date, awards, length) value ("Amelie",7.8,"2001-12-16 15:23:00",5, 122);

#Agregar un género a la tabla genres.
INSERT INTO movies_db.genres (created_at, name,ranking,active) value ("2021-12-16 16:00:00","Romance",13,1);
INSERT INTO movies_db.genres (created_at, name,ranking,active) value (NOW(),"Sci-Fi",14,1);

#Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE movies_db.movies SET genre_id = 13 WHERE id = 24;

#Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2
UPDATE movies_db.actors SET favorite_movie_id = 24 WHERE id = 3;

#Crear una tabla temporal copia de la tabla movies
CREATE TEMPORARY TABLE copy_movies (SELECT * FROM movies);

#Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
SET SQL_SAFE_UPDATES = 0;
SELECT * FROM copy_movies WHERE awards < 5;
DELETE FROM copy_movies WHERE awards<5;
SET SQL_SAFE_UPDATES = 1;

#Obtener la lista de todos los géneros que tengan al menos una película.
SELECT g.name, count(*) as cantidad_movies FROM genres g
INNER JOIN movies m ON g.id = m.genre_id
GROUP BY g.id HAVING count(*) >= 1;

#Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 

#Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

#¿Qué son los índices? ¿Para qué sirven?	
# Un índice SQL es una tabla de búsqueda rápida para poder encontrar los registros 
# que los usuarios necesitan buscar con mayor frecuencia y sirven para mejorar 
# sustancialmente los tiempos de respuesta

#Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX movies_name
ON movies_db.movies(id, title);

#Chequee que el índice fue creado correctamente.
show index from movies;
