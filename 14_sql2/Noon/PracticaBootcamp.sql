SELECT s.title, g.name
FROM series s
INNER JOIN genres g ON s.genre_id = g.id;

SELECT e.title, a.last_name, a.first_name
FROM episodes e
INNER JOIN actor_episode ae ON ae.episode_id = e.id
INNER JOIN actors a ON ae.actor_id = a.id
ORDER BY e.title;

SELECT ser.title, COUNT(sea.id)
FROM series ser
INNER JOIN seasons sea ON sea.serie_id = ser.id
GROUP BY ser.id;

SELECT g.name, COUNT(m.id) as "Total"
FROM genres g
INNER JOIN movies m ON m.genre_id = g.id
GROUP BY g.id
HAVING Total > 2;

SELECT DISTINCT a.last_name, a.first_name
FROM actors a
INNER JOIN actor_movie am ON am.actor_id = a.id
INNER JOIN movies m ON am.movie_id = m.id
WHERE m.title LIKE "%guerra de las galaxias%"
GROUP BY a.id
HAVING count(am.movie_id) >= (
	SELECT COUNT(*)
    FROM movies subm
    WHERE subm.title LIKE "%guerra de las galaxias%"
);