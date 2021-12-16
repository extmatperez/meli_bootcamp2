use movies_db;

SELECT s.title, g.name
FROM series s 
LEFT JOIN genres g ON s.genre_id = g.id;

SELECT e.title,a.first_name, a.last_name
FROM episodes e 
INNER JOIN actor_episode ae ON e.id = ae.episode_id
INNER JOIN actors a ON ae.actor_id = a.id;

SELECT s.title, COUNT(se.id) as "Nº Temporadas"
FROM series s 
INNER JOIN seasons se ON s.id = se.serie_id
GROUP BY s.id;

SELECT g.name, COUNT(m.id) as movies
FROM genres g 
LEFT JOIN movies m ON g.id = m.genre_id
GROUP BY g.id
HAVING movies > 2;

SELECT a.first_name, a.last_name
FROM actors a 
INNER JOIN actor_movie am ON a.id = am.actor_id
INNER JOIN movies m ON am.movie_id = m.id
WHERE m.title LIKE "%La Guerra de las galaxias%"
GROUP BY a.id;




