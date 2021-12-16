# Ejercicio 1: Mostrar el título y el nombre del género de todas las series.
SELECT series.title, genres.name
FROM series
JOIN genres ON series.genre_id = genres.id;

# Ejercicio 2: Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT episodes.title, actors.first_name, actors.last_name
FROM episodes
JOIN actor_episode ON actor_episode.episode_id = episodes.id
JOIN actors ON actors.id = actor_episode.actor_id;

# Ejercicio 3: Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT series.title, COUNT(seasons.id) as total
FROM series
JOIN seasons ON seasons.serie_id = series.id
GROUP BY (series.title);

# Ejercicio 4: Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT genres.name, COUNT(movies.id) as total
FROM genres
JOIN movies ON movies.genre_id = genres.id
GROUP BY (genres.name)
HAVING (total >= 3);

# Ejercicio 5: Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT DISTINCT actors.first_name, actors.last_name
FROM actors
JOIN actor_movie ON actor_movie.actor_id = actors.id
JOIN movies ON movies.id = actor_movie.movie_id
WHERE LOWER(movies.title) LIKE "la guerra de las galaxias%";
