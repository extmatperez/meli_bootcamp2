-- 1. Explicar el concepto de normalización y para que se utiliza.
-- Tener una base de datos estandarizada, se utiliza para tener los datos de manera organizada y sin dupilcaciones 
-- 2. Agregar una película a la tabla movies.
INSERT INTO MOVIES_DB.MOVIES(TITLE, RATING, AWARDS, RELEASE_DATE, LENGTH) VALUES ('Spiderman', 8.5, 4, '2021-12-16 0:00', 120);
-- 3. Agregar un género a la tabla genres.
INSERT INTO MOVIES_DB.GENRES(CREATED_AT, NAME, RANKING, ACTIVE) VALUES ('2021-12-16 0:00','Ficcion',13,1);
-- 4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.
UPDATE MOVIES_DB.MOVIES
SET GENRE_ID = 13
WHERE ID = 22 AND TITLE LIKE 'Spiderman';
-- 5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.
UPDATE MOVIES_DB.ACTORS
SET FAVORITE_MOVIE_ID = 22
WHERE ID = 47;
-- 6. Crear una tabla temporal copia de la tabla movies.
CREATE TEMPORARY TABLE MOVIES_DB.MOVIES_COPY 
SELECT * FROM MOVIES_DB.MOVIES;
-- 7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
SET SQL_SAFE_UPDATES = 0;
DELETE FROM MOVIES_DB.MOVIES_COPY WHERE AWARDS < 5;
SET SQL_SAFE_UPDATES = 1;
-- 8. Obtener la lista de todos los géneros que tengan al menos una película.
SELECT GEN.* FROM MOVIES_DB.GENRES AS GEN
JOIN  MOVIES_DB.MOVIES AS MOV ON GEN.ID = MOV.GENRE_ID
GROUP BY GEN.ID;
-- 9. Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
SELECT ACT.* FROM MOVIES_DB.ACTORS AS ACT
JOIN MOVIES_DB.MOVIES AS MOV ON MOV.ID = ACT.FAVORITE_MOVIE_ID
WHERE MOV.AWARDS > 3;
-- 10. Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
-- 11. ¿Qué son los índices? ¿Para qué sirven?
-- Mejoran el acceso a los datos, sirven para hacer busquedas de forma mas eficientes
-- 12. Crear un índice sobre el nombre en la tabla movies.
CREATE INDEX MOVIE_TITLE_INDEX ON MOVIES_DB.MOVIES(TITLE);
-- 13. Chequee que el índice fue creado correctamente.
SHOW INDEX FROM MOVIES_DB.MOVIES;



