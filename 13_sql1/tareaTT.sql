/*2. Mostrar todos los registros de la tabla de movies.*/
SELECT * FROM movies_db.movies;

/*3. Mostrar el nombre, apellido y rating de todos los actores.*/
SELECT first_name, last_name, rating FROM movies_db.actors;

/*4. Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla
como el campo estén en español*/
SELECT title AS titulo FROM movies_db.series;

/*5. Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5.*/
SELECT first_name, last_name FROM movies_db.actors
WHERE rating > 7.5;

/*6. Mostrar el título de las películas, el rating y los premios de las películas con un rating
mayor a 7.5 y con más de dos premios.*/
SELECT title, rating, awards FROM movies_db.movies
WHERE rating > 7.5 AND awards >2;

/*7. Mostrar el título de las películas y el rating ordenadas por rating en forma
ascendente.*/
SELECT title , rating FROM movies_db.movies
ORDER BY rating ASC;

/*9. Mostrar el top 5 de las películas con mayor rating.*/
SELECT title FROM movies_db.movies
ORDER BY rating DESC LIMIT 5;

/*10. Mostrar las top 5 a 10 de las películas con mayor rating.*/
SELECT title FROM movies_db.movies
ORDER BY rating DESC LIMIT 5 OFFSET 5;

/*11. Listar los primeros 10 actores (sería la página 1).*/
SELECT first_name, last_name FROM movies_db.actors
ORDER BY first_name ASC limit 10;

/*12. Luego usar offset para traer la página 3.*/
SELECT first_name, last_name FROM movies_db.actors
ORDER BY first_name ASC limit 10 offset 20;

/*13. Hacer lo mismo para la página 5.*/
SELECT first_name, last_name FROM movies_db.actors
ORDER BY first_name ASC limit 10 offset 40;

/*14. Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.*/
SELECT title , rating FROM movies_db.movies
WHERE title LIKE 'Toy Story%';

/*15. Mostrar a todos los actores cuyos nombres empiecen con Sam.*/
SELECT first_name, last_name FROM movies_db.actors
WHERE first_name LIKE 'Sam%';

/*16. Mostrar el título de las películas que salieron entre el 2004 y 2008.*/
SELECT title FROM movies_db.movies
WHERE YEAR(release_date) between '2004' AND '2008';

/*17. Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.*/
SELECT title FROM movies_db.movies
WHERE rating > 3 AND awards >1 AND (YEAR(release_date) between '1988' AND '2009')
ORDER BY rating;

/*18. Traer el top 3 a partir del registro 10 de la consulta anterior.*/
SELECT title FROM movies_db.movies
WHERE rating > 3 AND awards >1 AND (YEAR(release_date) between '1988' AND '2009')
ORDER BY rating LIMIT 3;