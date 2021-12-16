/*Actividad TT */
/*2*/
SELECT * FROM movies;
INSERT INTO movies (created_at,updated_at,title,rating,awards,release_date,length,genre_id)  
VALUES ("2018/10/10","2020/10/10","Pelicula Insertada 2",5,1,"2020/10/12",100,1);

/*3*/
SELECT * FROM genres;
INSERT INTO genres (created_at,updated_at,name,ranking,active) 
VALUES ("2018/10/10","2020/10/10","Miedo-Comica",20,1);

/*4*/
UPDATE movies SET genre_id = 13 WHERE id = 22;

/*5*/
SELECT * FROM actors;
UPDATE actors SET favorite_movie_id = 22 WHERE id < 3;

/*6*/
CREATE TEMPORARY TABLE tablaTempMovies (SELECT * FROM movies);
SELECT * FROM tablaTempMovies;

/*7*/
SET SQL_SAFE_UPDATES = 1;
DELETE FROM tablaTempMovies WHERE awards < 5;

/*8*/
SELECT * FROM movies m INNER JOIN genres g ON m.genre_id = g.id;

/*9*/