/*INSERT INTO `movies`
(`title`,
`rating`,
`release_date`,
`genre_id`)
VALUES
("Locademia de Policías",
7,
"1990-05-25",
1);*/


/*INSERT INTO `genres`
(`name`,
`ranking`,
`active`)
VALUES
("Western",
13,
1);*/

/*UPDATE `movies`
SET
`genre_id` = 13
WHERE `id` = 22;*/

/*UPDATE `actors`
SET
`favorite_movie_id` = 22
WHERE `id` = 47;*/


/*CREATE TEMPORARY TABLE peliculas (SELECT * FROM movies);

SELECT * FROM peliculas;*/

/*SET SQL_SAFE_UPDATES=0;
DELETE FROM peliculas WHERE awards < 5;
SET SQL_SAFE_UPDATES=1;
SELECT * FROM peliculas;*/

/*SELECT DISTINCT g.name
FROM genres g
INNER JOIN movies m ON m.genre_id = g.id*/

/*SELECT a.last_name, a.first_name
FROM actors a
INNER JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3;*/

/*ALTER TABLE movies ADD INDEX (title);*/