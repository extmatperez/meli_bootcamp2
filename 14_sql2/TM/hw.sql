1- SELECT series.title, genres.name FROM series INNER JOIN genres;
2- SELECT episodes.title, actors.last_name, actors.first_name from episodes INNER JOIN actors;
3 -SELECT series.title, count(seasons.serie_id) FROM series INNER JOIN seasons ON series.id = seasons.serie_id GROUP BY series.title;
4- SELECT genres.name, count(movies.genre_id) as "cantidad" FROM genres INNER JOIN movies ON genres.id = movies.genre_id GROUP BY genres.name HAVING cantidad >2;
5- SELECT DISTINCT first_name, last_name FROM actors INNER JOIN actor_movie INNER JOIN movies ON actor_movie.actor_id = actors.id AND actor_movie.movie_id = movies.id WHERE movies.title LIKE "La Guerra de las galaxias%";