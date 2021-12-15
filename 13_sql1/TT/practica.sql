-- 2. Mostrar todos los registros de la tabla de movies.
SELECT * FROM MOVIES;
-- 3. Mostrar el nombre, apellido y rating de todos los actores.
SELECT FIRST_NAME, LAST_NAME, RATING FROM ACTORS;
-- 4. Mostrar el título de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español
SELECT TITLE AS TITULO FROM SERIES AS SERIES;
-- 5. Mostrar el nombre y apellido de los actores cuyo rating sea mayor a 7.5
SELECT  FIRST_NAME, LAST_NAME FROM ACTORS WHERE RATING > 7.5;
-- 6. Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios.
SELECT TITLE, RATING, AWARDS FROM MOVIES WHERE RATING > 7.5 AND AWARDS > 2 ;
-- 7. Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
SELECT TITLE, RATING FROM MOVIES ORDER BY RATING DESC;
-- 8. Mostrar los títulos de las primeras tres películas en la base de datos.
SELECT TITLE FROM MOVIES LIMIT 3 OFFSET 0;
-- 9. Mostrar el top 5 de las películas con mayor rating
SELECT TITLE FROM MOVIES ORDER BY RATING DESC, TITLE ASC LIMIT 5 OFFSET 0;
-- 10. Mostrar las top 5 a 10 de las películas con mayor rating.
SELECT TITLE FROM MOVIES ORDER BY RATING DESC, TITLE ASC  LIMIT 5 OFFSET 5;
-- 11. Listar los primeros 10 actores (sería la página 1),
SELECT * FROM ACTORS  LIMIT 10 OFFSET 0;
-- 12. Luego usar offset para traer la página 3
SELECT * FROM ACTORS  LIMIT 10 OFFSET 20;
-- 13. Hacer lo mismo para la página 5
SELECT * FROM ACTORS  LIMIT 10 OFFSET 40;
-- 14. Mostrar el título y rating de todas las películas cuyo título sea de Toy Story.
SELECT TITLE, RATING FROM MOVIES WHERE TITLE LIKE 'Toy Story%';
-- 15. Mostrar a todos los actores cuyos nombres empiecen con Sam
SELECT * FROM ACTORS WHERE FIRST_NAME LIKE 'Sam%';
-- 16. Mostrar el título de las películas que salieron entre el 2004 y 2008.
SELECT TITLE FROM MOVIES WHERE RELEASE_DATE BETWEEN '2004/01/01' AND '2008/12/31'
-- 17. Traer el título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating
SELECT TITLE FROM MOVIES WHERE RATING > 3 AND AWARDS > 1 AND RELEASE_DATE BETWEEN '1988/01/01' AND '2009/12/31' ORDER BY RATING;
-- 18. Traer el top 3 a partir del registro 10 de la consulta anterior.
SELECT TITLE FROM MOVIES WHERE RATING > 3 AND AWARDS > 1 AND RELEASE_DATE BETWEEN '1988/01/01' AND '2009/12/31' ORDER BY RATING  LIMIT 3 OFFSET 10;



