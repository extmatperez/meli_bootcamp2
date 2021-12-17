/* Ejercicios prácticos */
/*1. Mostrar el título y el nombre del género de todas las series.*/
SELECT s.title, g.name FROM genres AS g JOIN series as s ON g.id = s.genre_id;

SELECT SER.TITLE, GEN.NAME FROM MOVIES_DB.SERIES AS SER JOIN MOVIES_DB.GENRES AS GEN ON SER.GENRE_ID =  GEN.ID;

SELECT s.title, g.name FROM series s INNER JOIN genres g;

/*2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos */
SELECT e.title, a.first_name, a.last_name 
FROM episodes AS e 
INNER JOIN actors AS a 
INNER JOIN actor_episode AS ae 
ON ae.actor_id = a.id AND ae.episode_id = e.id;

SELECT title, first_name, last_name FROM episodes AS ep
INNER JOIN actor_episode AS ac_ep
INNER JOIN  actors AS ac 
ON ep.id = ac_ep.episode_id AND ac.id = ac_ep.actor_id;

/*3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas. */
SELECT ser.title, (SELECT count(sea.number) FROM seasons AS sea 
WHERE ser.id = sea.serie_id) AS "total de temporadas" 
FROM series AS ser INNER JOIN seasons AS sea ON ser.id = sea.serie_id
GROUP BY ser.id;

SELECT SER.TITLE , COUNT(SER.ID) AS TEMPORADAS FROM MOVIES_DB.SEASONS AS SEA 
JOIN  MOVIES_DB.SERIES AS SER ON SEA.SERIE_ID = SER.ID
GROUP BY SER.ID;


/*4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3. */
SELECT name, (SELECT COUNT(id) FROM movies WHERE genre_id = gen.id) cantidad FROM genres AS gen
HAVING cantidad > 2;

/*5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan. */
SELECT DISTINCT a.first_name, a.last_name 
FROM movies AS m
INNER JOIN actors AS a 
INNER JOIN actor_movie AS am 
ON am.actor_id = a.id AND am.movie_id = m.id
WHERE m.title LIKE "La Guerra de las galaxias%";

SELECT DISTINCT ACT.FIRST_NAME AS NOMBRE, ACT.LAST_NAME AS APELLIDO FROM  MOVIES_DB.ACTOR_MOVIE AS ACT_MOV
JOIN MOVIES_DB.MOVIES AS MOV ON MOV.ID = ACT_MOV.MOVIE_ID AND MOV.TITLE LIKE 'La Guerra de las galaxias%'
JOIN MOVIES_DB.ACTORS AS ACT ON ACT.ID = ACT_MOV.ACTOR_ID;