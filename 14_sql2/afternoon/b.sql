#1. Explicar el concepto de normalización y para que se utiliza.

La normalización es un proceso de estandarización y validación de datos que consiste en eliminar las redundancias o inconsistencias, completando datos mediante una serie de reglas que actualizan la información, protegiendo su integridad y favoreciendo la interpretación, para que así sea más simple de consultar y más eficiente para quien la gestiona.
 
#2. Agregar una película a la tabla movies.

INSERT INTO movies
(title, rating, awards, release_date)
VALUES
('Super Bootcamp: Meli', 10, 99, '2021-11-15')

#3. Agregar un género a la tabla genres.

INSERT INTO genres
(name, ranking, active)
VALUES
('Kahoot!', 13, 1)

#4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.

UPDATE movies m
SET genre_id = 13
WHERE m.id = 22

5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.

UPDATE actors a
SET favorite_movie_id = 22
WHERE a.id = 4

#6. Crear una tabla temporal copia de la tabla movies.

CREATE TEMPORARY TABLE tabla_temporal
(SELECT *
FROM movies);

extra, consultar una tabla temporal (es igual que cualquier otra jaja):

SELECT *
FROM tabla_temporal;

var la estructura de la tabla temporal(es igual que cualquier otra jaja):

DESC tabla_temporal2;

#7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

DELETE
FROM tabla_temporal
WHERE awards < 5;

#8. Obtener la lista de todos los géneros que tengan al menos una
película.

SELECT distinct g.name
FROM movies m
INNER JOIN genres g ON m.genre_id = g.id

extra, muetra la cantidad de peliculas que tiene el genero:

SELECT g.name, count(g.name) total
FROM movies m
INNER JOIN genres g ON m.genre_id = g.id
GROUP BY g.name
HAVING total >=1

#9. Obtener la lista de actores cuya película favorita haya ganado más
de 3 awards.

SELECT a.first_name, a.last_name
FROM actors a
INNER JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3

EXTRA, mostrando la pelicula y cantidad de premios

SELECT a.first_name nombre, a.last_name apellido, m.title titulo, m.awards 'cantidad de premios'
FROM actors a
INNER JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3 

#10. Utilizar el explain plan para analizar las consultas del Ej.6 y 7.

¿?

#11. ¿Qué son los índices? ¿Para qué sirven?

● Son un mecanismo para optimizar consultas en SQL.
● Mejoran sustancialmente los tiempos de respuesta en Queries complejas.
● Mejoran el acceso a los datos al proporcionar una ruta más directa a los registros.
● Evitan realizar escaneos (barridas) completas o lineales de los datos en una tabla.

#12. Crear un índice sobre el nombre en la tabla movies.

indice unico: (no puede crearse sobre una columna que ya tenga datos duplicados)

CREATE UNIQUE INDEX moviesNameIndex
ON movies(title)

variante: 

ALTER TABLE movies
ADD UNIQUE INDEX moviesIdIndex(id)

indice que admite repeticiones:

CREATE INDEX moviesRatingIndex
ON movies(rating)

variante:

ALTER TABLE movies
ADD INDEX moviesTitleIndex(title)

#13. Chequee que el índice fue creado correctamente.

SHOW INDEX FROM movies;