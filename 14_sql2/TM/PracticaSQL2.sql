use movies_db;
##1.¿A qué se denomina JOIN en una base de datos?
##A la intersección de datos entre dos conjuntos, en el caso de SQL entre dos tablas, 
##asociadas por la llave foránea.

##2.Nombre y explique 2 tipos de JOIN.
##Left join: toma los datos de la tabla A y sus coincidencias de la tabla B, 
##sin importar que no tenga relación con algún registro de la otra tabla.

##Inner Join: Toda la intersección directa entre dos conjuntos de datos donde 
##debe estar relacionado con la llave foránea.

##3.¿Para qué se utiliza el GROUP BY?
##Para realizar agrupaciones de una consulta por uno o varios campos.

##4.¿Para qué se utiliza el HAVING?
##Este se utiliza para filtrar  una consulta después de un group by.

##5. Dado lo siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:

##Inner Join

##Left Join

##6. Escriba una consulta genérica por cada uno de los diagramas a continuación:

##Rigth Join:
SELECT a.first_name,a.last_name,m.title from movies m 
RIGHT JOIN actors a ON a.favorite_movie_id = m.id;

##Full Join:
SELECT m.title, g.name FROM movies m
LEFT JOIN genres g ON m.id = g.id
UNION
SELECT m.title, g.name FROM movies m
RIGHT JOIN genres g ON m.id = g.id and g.id is null;

## SEGUNDA PARTE
##1. Mostrar el título y el nombre del género de todas las series.
SELECT s.title, g.name FROM series s
INNER JOIN genres g ON s.id = g.id;

##2. Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
SELECT e.title FROM episodes e
INNER JOIN actor_episode ae on ae.episode_id=e.id
INNER JOIN actors a on a.id=ae.actor_id;

##3. Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
SELECT sr.title, se.numero FROM series sr
INNER JOIN (select serie_id, max(number) numero from seasons group by serie_id )se on se.serie_id=sr.id;

##4. Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
SELECT g.name,count(m.id) cantidad FROM genres g
INNER JOIN movies m on g.id =m.genre_id
GROUP BY g.name
HAVING count(m.id)>=3;

##5. Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
SELECT a.first_name, a.last_name FROM actors a
INNER JOIN  actor_movie am on am.actor_id=a.id
INNER JOIN movies m on m.id=am.movie_id and m.title like '%guerra de las galaxias%'
GROUP BY a.first_name, a.last_name;