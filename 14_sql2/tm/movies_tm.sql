USE movies_db;

# Punto 1
SELECT s.title, g.name FROM series s INNER JOIN genres g ON s.genre_id = g.id;

#Punto 2
SELECT e.title, a.first_name, a.last_name
FROM episodes e
INNER JOIN actor_episode ae ON e.id = ae.episode_id
INNER JOIN actors a ON ae.actor_id = a.id;

#Punto 3
SELECT s.title, count(se.id) AS "season amount"
FROM series s
INNER JOIN seasons se ON s.id = se.serie_id
GROUP BY s.id;

#Punto 4
SELECT g.name, count(m.id) AS amount
FROM genres g
INNER JOIN movies m ON g.id = m.genre_id
GROUP BY g.id
HAVING amount >= 3;

#Punto 5
SELECT DISTINCT a.first_name, a.last_name
FROM actors a
INNER JOIN actor_movie am ON a.id = am.actor_id
INNER JOIN movies m ON am.movie_id = m.id
WHERE m.id IN (SELECT m.id FROM movies m WHERE m.title LIKE "%Guerra de las Galaxias%");