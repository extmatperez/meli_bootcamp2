/*Explicar el concepto de normalización y para que se utiliza.*/
/* es un proceso de estandarización y validación de datos que consiste en eliminar
las redundancias o inconsistencias, completando datos mediante una serie de reglas que actualizan
la información, protegiendo su integridad y favoreciendo la interpretación, para que así sea más
simple de consultar y más eficiente para quien la gestiona.*/

/*Agregar una película a la tabla movies.*/
insert into movies(title,rating,awards,release_date,genre_id) 
values ("PeliNueva",8,0,"2021-12-16",8);

/*Agregar un género a la tabla genres.*/
insert into genres(created_at,name,ranking,active) 
values ("2021-12-16","NuevoGenero",13,1);


/*Asociar a la película del Ej 2. con el género creado en el Ej. 3.*/
update movies set genre_id=13 where id=22;

/*Modificar la tabla actors para que al menos un actor 
tenga como favorita la película agregada en el Ej.2.*/
update actors set favorite_movie_id=22 where id < 25;

/*Crear una tabla temporal copia de la tabla movies.*/
create temporary table moviestemporal select * from movies;
select * from moviestemporal;

/*Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.*/
delete from moviestemporal where awards >= 5;
select * from moviestemporal;

/*Obtener la lista de todos los géneros que tengan al menos una película.*/
select ge.name,count(ge.id) as cantidad from genres ge 
join movies mo on mo.genre_id = ge.id
group by ge.id 
having cantidad >= 1;

/*Obtener la lista de actores cuya película favorita haya ganado más de 3 awards. */
select distinct(ac.id),ac.first_name from actors ac
join movies mo on ac.favorite_movie_id=mo.id
where mo.awards > 3;

/*Utilizar el explain plan para analizar las consultas del Ej.6 y 7.*/
explain  select * from movies;
explain delete from moviestemporal where awards >= 5;

/*¿Qué son los índices? ¿Para qué sirven?*/
/*● Son un mecanismo para optimizar consultas en SQL.
● Mejoran sustancialmente los tiempos de respuesta en Queries complejas.
● Mejoran el acceso a los datos al proporcionar una ruta más directa a los registros.
● Evitan realizar escaneos (barridas) completas o lineales de los datos en una tabla*/

/*Crear un índice sobre el nombre en la tabla movies.*/
SHOW INDEX FROM movies;
CREATE INDEX movies_idx ON movies (release_date);

/*Chequee que el índice fue creado correctamente.*/
SHOW INDEX FROM movies;
