-- Parte 1
-- 1. ¿A qué se denomina JOIN en una base de datos?
-- Es la union de dos tablas apartir de un dato comun 
-- 2. Nombre y explique 2 tipos de JOIN.
-- Inner Join: La union de datos de tabla a y tabla b que tengan coincidencias en ambas tablas
-- Left Join: La union de todos los datos de la primera tabla combinando con los existentes de segunda tabla 
-- 3. ¿Para qué se utiliza el GROUP BY?
-- Agrupa un conjunto de datos que se encuentren repetidos
-- 4. ¿Para qué se utiliza el HAVING?
-- Es un filtro que se aplica a los datos agrupados
-- 5. Dado lo siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
-- A. Inner Join - B. Left Join
-- 6. Escriba una consulta genérica por cada uno de los diagramas a continuación:
-- SELECT * FROM TABLA_A RIGHT JOIN TABLA_B ON TABLA_A.ID = TABLA_B.ID_TABLAB
-- SELECT * FROM TABLA_A FULL OUTER JOIN TABLA_B ON TABLA_A.ID = TABLA_B.ID_TABLAB
 
SELECT * FROM MOVIES_DB.MOVIES;
-- Parte 2
-- 1. Mostrar el título y el nombre del género de todas las series.
SELECT SER.TITLE, GEN.NAME FROM MOVIES_DB.SERIES AS SER JOIN MOVIES_DB.GENRES AS GEN ON SER.GENRE_ID =  GEN.ID;
-- 2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT EPI.TITLE, ACT.FIRST_NAME, ACT.LAST_NAME FROM MOVIES_DB.ACTOR_EPISODE AS ACT_EPI
JOIN MOVIES_DB.ACTORS AS ACT ON ACT.ID = ACT_EPI.ACTOR_ID
JOIN MOVIES_DB.EPISODES AS EPI ON EPI.ID = ACT_EPI.EPISODE_ID ;


-- 3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT SER.TITLE , COUNT(SER.ID) AS TEMPORADAS FROM MOVIES_DB.SEASONS AS SEA 
JOIN  MOVIES_DB.SERIES AS SER ON SEA.SERIE_ID = SER.ID
GROUP BY SER.ID;


-- 4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT GEN.NAME , COUNT(GEN.ID) AS PELICULAS FROM MOVIES_DB.GENRES AS GEN
JOIN MOVIES_DB.MOVIES AS MOV ON MOV.GENRE_ID = GEN.ID
GROUP BY GEN.ID
HAVING PELICULAS > 2;

-- 5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.

SELECT DISTINCT ACT.FIRST_NAME AS NOMBRE, ACT.LAST_NAME AS APELLIDO FROM  MOVIES_DB.ACTOR_MOVIE AS ACT_MOV
JOIN MOVIES_DB.MOVIES AS MOV ON MOV.ID = ACT_MOV.MOVIE_ID AND MOV.TITLE LIKE 'La Guerra de las galaxias%'
JOIN MOVIES_DB.ACTORS AS ACT ON ACT.ID = ACT_MOV.ACTOR_ID;

SELECT ACT.FIRST_NAME AS NOMBRE, ACT.LAST_NAME AS APELLIDO FROM  MOVIES_DB.ACTOR_MOVIE AS ACT_MOV
JOIN MOVIES_DB.MOVIES AS MOV ON MOV.ID = ACT_MOV.MOVIE_ID
JOIN MOVIES_DB.ACTORS AS ACT ON ACT.ID = ACT_MOV.ACTOR_ID
WHERE MOV.TITLE LIKE 'La Guerra de las galaxias%'
GROUP BY ACT.FIRST_NAME, ACT.LAST_NAME;



