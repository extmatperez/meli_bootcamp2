# PRIMERA PARTE

#1. ¿A qué se denomina JOIN en una base de datos?
	/* el JOIN es una sentencia SQL que permite combinar registros entre multiples tabla  */
    
#2. ¿A qué se denomina JOIN en una base de datos?
	/* 
		el INNER JOIN combina y trae los registros que tengan datos comunes entre dos tabla  mientras que un 
		LEFT JOIN trae todos los datos de la primera table y si tiene datos la segunda los trae, si no tienen
        coincidencia trae el campo como null
	*/
    
#3. ¿Para qué se utiliza el GROUP BY?
	/*agrupar los datos de las consultas por campos determinado luego de la sentencia */

#4. ¿Para qué se utiliza el HAVING?
	/*
		HAVING tiene la misma utilidad que el WHERE con la diferencia que HAVING se usa para datos que fueron 
		agrupados mientras que el WHERE se usa para datos comunes sin agrupamiento
	*/
    
#5. Dado lo siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
	/* INNER JOIN y LEFT JOIN*/

#6 Escriba una consulta genérica por cada uno de los diagramas a continuación:

	/*
		SELECT * 
        FROM `tabla_1`
		RIGHT JOIN `tabla_2` ON `tabla_1.id_tabla2` = `tabla_2.id`
        
        SELECT * 
        FROM `tabla_1`
		FULL JOIN `tabla_2` ON `tabla_1.id_tabla2` = `tabla_2.id`
    */
    
# SEGUNDA PARTE

# 1. Mostrar el título y el nombre del género de todas las series.
SELECT S.*, G.name  
FROM series AS S
INNER JOIN genres AS G ON S.genre_id = G.id;

#2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

SELECT E.title, CONCAT(A.first_name,' ', A.last_name) AS actor
FROM actor_episode AS AE
INNER JOIN episodes AS E ON AE.episode_id = E.id
INNER JOIN actors AS A ON AE.actor_id = A.id;

#3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.

SELECT series.title AS Serie , COUNT(seasons.title) AS 'Total de Temporadas' 
FROM seasons
INNER JOIN series ON seasons.serie_id = series.id
GROUP BY serie_id;

#4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

SELECT G.name AS Genero, COUNT(G.name) AS N_peliculas
FROM genres AS G
INNER JOIN movies AS M ON G.id = M.genre_id
GROUP BY genre_id
HAVING N_peliculas >= 3;

# Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.

SELECT distinct CONCAT(A.first_name,' ' ,A.last_name) AS Actor
FROM actor_movie AS AC
INNER JOIN actors AS A ON AC.actor_id = A.id 
INNER JOIN movies AS M ON AC.movie_id = M.id
WHERE title LIKE '%guerra%';























