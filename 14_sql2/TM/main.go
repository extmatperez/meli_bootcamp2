//¿A qué se denomina JOIN en una base de datos?
//Consiste en combinar datos de una tabla con datos de otra tabla mediante una o varias condiciones en comun

//Nombre y explique 2 tipos de JOIN.
//INNER JOIN : Es la interseccion entre dos tablas.
//LEFT JOIN: : es una interseccion pero aca le damos prioridad a la tabla de la izquierda, y
//buscamos en la tabla de la derecha, se muestra si o si lo de la izquierdada

//¿Para qué se utiliza el GROUP BY?
//Sirve para agrupar resultados segun las columnas indicadas, genera un solo registro
//por cada grupo

//¿Para qué se utiliza el HAVING?
//Es muy parecido al WHERE,pero HAVING aplica un filtro a los grupos obtenidos por el GROUP BY

//5.A) INNER JOIN
//5.B) LEFT JOIN

//6.A) SELECT * FROM table1 RIGHT JOIN table2 ON table1.id = table2.movie_id
//6.B) SELECT * FROM table1 FULL JOIN table2 ON table1.commond_field = table2.commond_field

//EJE: 1)
//Mostrar el título y el nombre del género de todas las series.
//SELECT series.title,genres.name
//FROM movies_db.series series
//INNER JOIN movies_db.genres genres
//ON series.genre_id = genres.id;

//EJE: 2)
//Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
//SELECT episodes.title,actors.first_name,actors.last_name
//FROM movies_db.episodes episodes
//INNER JOIN movies_db.actor_episode actor_episode ON episodes.id = actor_episode.id
//INNER JOIN movies_db.actors actors ON actor_episode.actor_id = actors.id;

//EJE: 3)
//Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
//SELECT series.title as titulo, COUNT(seasons.serie_id) as temporadas
//FROM movies_db.series series
//INNER JOIN movies_db.seasons seasons ON series.id = seasons.serie_id
//GROUP BY serie_id;

//EJE: 4)
//Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno,
// siempre que sea mayor o igual a 3.
//
