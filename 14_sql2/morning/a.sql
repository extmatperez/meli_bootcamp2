PRIMERA PARTE


1. ¿A qué se denomina JOIN en una base de datos?

    JOIN es la instruccion utilizada para combinar la informacion de dos tablas utilizando uno o mas campos de comparacion

2. Nombre y explique 2 tipos de JOIN.

    1. INNER JOIN: Es el join por defecto y puede escribirse solamente JOIN, este caso se utiliza para generar la interseccion entre dos tablas.
    2. LEFT JOIN: Este join, genera la interseccion entre dos tablas mas todos los datos de la tabla de la izquierda aunque no tenga match con la tabla de la derecha.

3. ¿Para qué se utiliza el GROUP BY?

    El GROUP BY es utilizado para agrupar los datos de una tabla en una sola fila.

4. ¿Para qué se utiliza el HAVING?

    Es como un WHERE pero para se aplica a los grupos obtenidos por el GROUP BY

5. Dado lo siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:

    1. caso de la izquierda: corresponde a un INNER JOIN
    2. caso de la derecha: corresponde a un LEFT JOIN

6. Escriba una consulta genérica por cada uno de los diagramas a continuación:

    1.  SELECT *
        FROM movies 
        RIHGH JOIN genres
        ON genre_id = genres.id

    2.  SELECT *
        FROM movies
        FULL JOIN genres


SEGUNDA PARTE


1. Mostrar el título y el nombre del género de todas las series.

SELECT s.title, g.name
FROM series s
INNER JOIN genres g ON s.genre_id = g.id

2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

SELECT e.title, a.first_name, a.last_name
FROM episodes e
INNER JOIN actor_episode ae ON e.id = ae.episode_id
INNER JOIN actors a ON a.id = ae.actor_id

3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.

SELECT s.title, count(s.title)
FROM series s
INNER JOIN seasons se ON s.id = se.serie_id
GROUP BY s.title

4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

SELECT g.name as 'Genre Name', count(g.name) as Total
FROM movies m
INNER JOIN genres g ON m.genre_id = g.id
GROUP BY g.name
HAVING total >= 3

5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.

SELECT distinct a.first_name, a.last_name
FROM actors a
INNER JOIN actor_movie am ON a.id = am.actor_id
INNER JOIN (
	SELECT m.id, m.title
    FROM movies m
    WHERE m.title LIKE 'La Guerra de las galaxias%'
) m ON am.movie_id = m.id
