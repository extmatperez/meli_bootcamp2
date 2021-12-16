### PRACTICA TT SQL 3 ###
set SQL_SAFE_UPDATE = 0;
use movies_db;
## 1. Explicar el concepto de normalización y para que se utiliza.
## Es un mecanismo que se utiliza para esquematizar la base de datos de una forma adecuada
## eliminando la redundancia de datos y mejorando la integridad. LLevando a cabo las 4 formas 
## normales se mejora la eficiencia de las consultas.

## 2. Agregar una película a la tabla movies.
INSERT INTO movies (title,rating,awards,release_date)values ("Gol",8.0,2,"1996-09-16");

## 3. Agregar un género a la tabla genres.
INSERT INTO genres (created_at,name,ranking,active) values ('2013-07-03 22:00:00', 'Deportes',13, 1);

## 4.Asociar a la película del Ej 2. con el género creado en el Ej. 3.
Update movies 
set genre_id=15
where id = 22;

## 5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE actors
set favorite_movie_id= 15
where actors.id=47;

## 6. Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE movies_copia
(select * from movies);
select * from movies_copia;

## 7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
DELETE FROM movies_copia
WHERE awards>4;

## 8. Obtener la lista de todos los géneros que tengan al menos una película.
SELECT g.name, count(m.id) cantidad FROM genres g
INNER JOIN movies m on g.id=m.genre_id
GROUP BY g.name
HAVING count(m.id)>0 ;

## 9. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
SELECT a.first_name,a.last_name,m.awards FROM actors a
INNER JOIN movies m on a.favorite_movie_id=m.id and m.awards>2;

## 10. Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
##OK

## 11. ¿Qué son los índices? ¿Para qué sirven?
## Es un mecanismo que permite ubicar información en una tabla, esto mejora los tiempos de respuesta
## al ser un valor único y en ocaciones es secuencial.  

## 12. Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX name_idx ON movies(title);

## 13. Chequee que el índice fue creado correctamente.
SHOW INDEX FROM movies;
