# 2. Agregar una película a la tabla movies.
/*
	INSERT INTO `movies_db`.`movies`(`title`,`rating`,`awards`,`release_date`,`length`,`genre_id`) VALUES('interestelar',10,10,'2017-01-01',120,5);
*/
# 3.Agregar un género a la tabla genres.alter

# 4. Asociar a la película del Ej 2. con el género creado en el Ej. 3.

# 5. Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el Ej.2.

# 6. Crear una tabla temporal copia de la tabla movies.
	 /* CREATE TEMPORARY TABLE movies_2 (SELECT  * FROM movies); */
# 7. Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.
    /* DELETE FROM movies_2 WHERE awards < 5; */
    
# 8. Obtener la lista de todos los géneros que tengan al menos una película.
	SELECT genre.name, COUNT(*) as gen 
	FROM genres genre, movies movie
	WHERE genre.id = movie.genre_id
	GROUP BY movie.genre_id
	HAVING gen >= 1;
#Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. 
# Utilizar el explain plan para analizar las consultas del Ej.6 y 7.
	
# ¿Qué son los índices? ¿Para qué sirven?
# Crear un índice sobre el nombre en la tabla movies.
# Chequee que el índice fue creado correctamente.