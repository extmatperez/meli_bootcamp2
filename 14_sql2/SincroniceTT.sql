select * from movies_db.movies;

INSERT INTO movies_db.movies (title, rating, release_date, awards) value ("insert prueba",9.1,"2021-12-16 15:23:00",5);

UPDATE movies_db.movies SET rating=3.4 WHERE id >= 22;

DELETE FROM movies_db.movies WHERE id >= 23;

use movies_db;

#Creacion de tabla temporal
CREATE TEMPORARY TABLE moviesFilter (SELECT * FROM movies WHERE rating > 8);

SELECT * FROM moviesFilter;

CREATE view viewOfMovie AS (select * from movies where awards > 4);

select * from viewOfMovie;

select max(rating) from viewOfMovie
union
select min(rating) from viewOfMovie
union
select avg(rating) from viewOfMovie;

show index from users;