USE movies_db;
/*1. Mostrar el título y el nombre del género de todas las series.*/
SELECT genres.name as Genero, series.title as Titulo 
FROM movies_db.genres genres
INNER JOIN movies_db.series series
ON series.genre_id = genres.id;
/*2. Mostrar el título de los episodios, el nombre y apellido de los actores que
trabajan en cada uno de ellos.*/
SELECT ac.first_name,ac.last_name,ep.title 
from actor_episode ae 
join actors ac on ae.actor_id = ac.id 
join episodes ep on ep.id= ae.episode_id;

/*3. Mostrar el título de todas las series y el total de temporadas que tiene
cada una de ellas.*/
select se.title, max(sea.number) as "Total de temporadas" 
from series se 
join seasons sea on sea.serie_id = se.id 
group by title;

/*4. Mostrar el nombre de todos los géneros y la cantidad total de películas
por cada uno, siempre que sea mayor o igual a 3.*/
SELECT ge.name, count(ge.name) as cont 
FROM genres ge 
JOIN movies mo ON mo.genre_id = ge.id 
GROUP BY ge.name
HAVING cont >= 3;

/*5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas
las películas de la guerra de las galaxias y que estos no se repitan.*/
SELECT DISTINCT first_name, last_name
FROM movies_db.actors
JOIN movies_db.actor_movie ON actor_movie.actor_id = actors.id
JOIN movies_db.movies ON movies.id = actor_movie.movie_id
WHERE LOWER(movies.title) LIKE "%guerra de las galaxias%";