/*C2 SQL 2 TT
PARTE 1
1- ¿A qué se denomina JOIN en una base de datos?
2- Nombre y explique 2 tipos de JOIN.
3- ¿Para qué se utiliza el GROUP BY?
4- ¿Para qué se utiliza el HAVING?
5- Dado lo siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
6- Escriba una consulta genérica por cada uno de los diagramas a continuación:
*/

/*PARTE 2*/
-- Mostrar el título y el nombre del género de todas las series.
SELECT series.title, genres.name
FROM series
LEFT JOIN genres
ON series.genre_id = genres.id;


-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT episodes.title, actors.first_name, actors.last_name
FROM episodes
LEFT JOIN actor_episode
ON episodes.id = actor_episode.episode_id
LEFT JOIN actors
ON actor_episode.actor_id = actors.id;

-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT series.id,series.title, count(seasons.id)
FROM series LEFT JOIN seasons
ON series.id = seasons.serie_id
GROUP BY series.id;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT genres.name, count(movies.id)
FROM genres
LEFT JOIN movies
ON genres.id = movies.genre_id
GROUP BY genres.id
HAVING count(movies.id)>=3;

-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
-- >> Actuaron en al menos una película:
SELECT actors.first_name, actors.last_name
FROM(
	SELECT distinct actors.id, actors.first_name, actors.last_name
	FROM actors
	LEFT JOIN actor_movie
	ON actors.id = actor_movie.actor_id
	LEFT JOIN movies
	ON actor_movie.movie_id = movies.id
	WHERE movies.title LIKE "%guerra de las galaxias%"
) as actors;

-- >> Actuaron en TODAS las películas:
SELECT a.first_name, a.last_name
FROM actors a
WHERE EXISTS (
	SELECT *
    FROM movies m
    WHERE m.title LIKE "%guerra de las galaxias%"
)AND NOT EXISTS (
	SELECT *
    FROM movies m
    WHERE m.title LIKE "%guerra de las galaxias%"
    AND NOT EXISTS (
		SELECT *
        FROM actor_movie am
        WHERE am.actor_id=a.id AND am.movie_id=m.id
	)
);





