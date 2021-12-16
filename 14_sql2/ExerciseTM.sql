#Mostrar el título y el nombre del género de todas las series.
SELECT title, g.name
FROM series s
INNER JOIN genres g
ON s.genre_id = g.id;

#Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT epi.title, a.first_name, a.last_name
FROM episodes epi
INNER JOIN actor_episode act
ON epi.id = act.episode_id
JOIN actors a
ON act.actor_id = a.id;

#Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT SER.title AS "Titulo serie",
count(SER.id) AS "Numero de Temporadas"
FROM seasons SEA
JOIN series SER
ON SER.id = SEA.serie_id
GROUP BY SER.id;

#Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT GEN.name, count(MOV.id) AS cantidad
FROM movies MOV
JOIN genres GEN
ON GEN.id = MOV.genre_id
GROUP BY MOV.genre_id
HAVING cantidad > 2;

#Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de
# la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT actors.first_name, actors.last_name
FROM actors
JOIN actor_movie ON actor_movie.actor_id = actors.id
JOIN movies ON movies.id = actor_movie.movie_id
WHERE LOWER(movies.title) LIKE "la guerra de las galaxias%";