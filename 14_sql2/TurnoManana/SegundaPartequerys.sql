/*Mostrar el título y el nombre del género de todas las series.*/
SELECT s.title,g.name FROM series s 
JOIN genres g on s.genre_id = g.id;

/*Mostrar el título de los episodios, 
el nombre y apellido de los actores que trabajan en cada uno de ellos.*/

select e.title,a.first_name,a.last_name from actor_episode ae
join actors a on ae.actor_id = a.id
join episodes e on ae.episode_id = e.id;

/*Mostrar el título de todas las series y el total de temporadas 
que tiene cada una de ellas.*/

select s.title, count(s.id) as "Total Temporadas" 
from series s
join seasons se on s.id = se.serie_id
group by s.id;

/*Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, 
siempre que sea mayor o igual a 3.*/

select ge.name, count(mo.id) as "total peliculas" from movies mo 
join genres ge on mo.genre_id = ge.id
group by  ge.id
having total > 3;

/*Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas 
de la guerra de las galaxias y que estos no se repitan*/

select a.first_name,a.last_name from actor_movie am
join actors a on a.id = am.actor_id
join movies m on m.id= am.movie_id
where m.title like "%La Guerra de las galaxias%" 
group by a.id
