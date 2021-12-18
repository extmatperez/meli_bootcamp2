1- es un proceso de estandarizacion y validacion de datos que elimina redundancias o inconsistencias y completa los datos mediante 
una serie de reglas que actualizan la informacion protegiendo su integridad y favoreciendo la interpretacion. para simplificar 
consultas y hacer la gestion mas eficientes
2- INSERT INTO movies (title,rating,awards,release_date) VALUES ("rocky 4", 10.0, 100, "1985-01-01 20:40:00")
3- INSERT INTO genres(name,ranking) VALUES ("river", 200);
4- UPDATE movies SET genre_id = 14 WHERE title = "rocky 4";
5- UPDATE actors SET favorite_movie_id = 22 WHERE id=3;
6- CREATE TEMPORARY TABLE moviesCopyy (SELECT * FROM movies);
7- DELETE FROM moviesCopyy WHERE awards < 5;
8- SELECT genres.name, COUNT(movies.id) as peliculas FROM genres INNER JOIN movies on  genres.id = movies.genre_id 
GROUP BY genres.id HAVING peliculas >1;
9- SELECT  a.*, movies.awards FROM actors a INNER JOIN movies ON a.favorite_movie_id = movies.id 
GROUP BY a.id HAVING movies.awards >3;
12- CREATE INDEX movies_title ON movies (title);
13- SHOW INDEX FROM movies;

