-- 1 titulo y genero de todas las series 
select se.title, ge.name 
from series se left join genres ge on se.genre_id = ge.id;

-- 2 
-- Mostrar el título de los episodios, el nombre y apellido 
-- de los actores que trabajan en cada uno de ellos.


-- 3


-- 4 Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
select serie_id, count(*) as cantidad_seasons
from seasons 
group by serie_id;



-- 5 Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.



-- 6 Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
