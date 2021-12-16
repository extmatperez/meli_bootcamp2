select * from actors;

SELECT avg(rating) as "Promedio Rating" FROM movies;

#Agrupacion del promedio de ranking de cada genero de la tabla movies
SELECT avg(rating) as "Promedio Rating", genre_id 
FROM movies 
GROUP BY genre_id;

#Agrupacion del promedio de ranking de cada genero de la tabla movies 
#Con promedios mayores que 7 {HAVING}
SELECT avg(rating) as promedio, genre_id 
FROM movies 
GROUP BY genre_id
HAVING promedio > 7;

#Obtener todos los promedios de los rating de los generos que sea mayor
#que el promedio de los generos totales.
SELECT avg(rating) as promedio, genre_id 
FROM movies 
GROUP BY genre_id
HAVING promedio > (SELECT AVG(rating) FROM movies);

#Traer el promedio de todos los movies agrupados por genre_id sin usar HAVING
SELECT avg(rating) as promedio, genre_id 
FROM movies 
GROUP BY genre_id
HAVING promedio > (SELECT AVG(rating) FROM movies);
SELECT DISTINCT genre_id FROM movies WHERE genre_id > 0;
