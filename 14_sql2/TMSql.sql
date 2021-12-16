/* SQL II */

/*1*/
SELECT ge.name , s.title FROM series s INNER JOIN genres ge ON s.genre_id = ge.id;

/*2*/
SELECT e.title ,ac.first_name,ac.last_name 
FROM actors ac 
INNER JOIN actor_episode ae 
ON ae.actor_id = ac.id
INNER JOIN episodes e
ON ae.episode_id = e.id;

/*3*/
SELECT s.title, MAX(ss.number) FROM series s INNER JOIN seasons ss ON ss.serie_id = s.id GROUP BY s.title;

/*4*/
SELECT g.name, COUNT(g.name)  FROM genres g  INNER JOIN movies m ON m.genre_id = g.id GROUP BY g.name;
SELECT * FROM movies ;
SELECT * FROM genres ;

/*5*/
SELECT a.first_name, a.last_name FROM actors a 
INNER JOIN actor_movie am ON am.actor_id = a.id 
INNER JOIN movies m ON am.movie_id = m.id
WHERE m.title LIKE '%Guerra de las galaxias%'