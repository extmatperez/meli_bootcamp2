use movies_db;
-- 1. Se trata de una serie de reglas las cuales permiten una estandarización en los tados, evitando redundancias o inconsistencias en las tablas.

INSERT INTO movies (title,rating,awards,release_date,length,genre_id)
VALUES ("Spiderman: No Way Home", 9.8, 0, "2021-12-15 00:00:00", 180, 5);

INSERT INTO genres (created_at,name, ranking, active)
VALUES("2021-12-16 00:00:00","Romance", 13, 1);

UPDATE movies SET genre_id = 13
WHERE title LIKE "%Spiderman%";

UPDATE actors SET favorite_movie_id = 22
WHERE Id = 13;

CREATE TEMPORARY TABLE moviesCopy (SELECT * FROM movies);

DELETE FROM moviesCopy WHERE awards < 5;

SELECT g.*, COUNT(m.id) as peliculas
FROM genres g
LEFT JOIN movies m ON g.id = m.genre_id
GROUP BY g.id
HAVING peliculas > 1;

SELECT a.*, m.awards
FROM actors a 
LEFT JOIN movies m ON m.id = a.favorite_movie_id
WHERE m.awards > 3;

-- Permiten indexar columnas específicas en una tabla con el fin de agilizar consultas y facilitar el acceso a datos.


CREATE INDEX movies_title ON movies (title);

SHOW INDEX FROM movies;




