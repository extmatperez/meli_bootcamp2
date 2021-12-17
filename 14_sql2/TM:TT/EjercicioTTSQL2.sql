# Normalizacion
# La normalizacion es la practica sobre una base de datos que consiste en la estandarizacion y validacion de datos, eliminando las redundancias y las inconsistencias,
# con una serie de practicas (definidas en 4 formas normales, o FN) protegiendo la identidad y favoreciendo la interpretacion de las mismas.

# RNF:
# 1FN: Crea registros independientes, por ejemplo, un campo para un telefono y no dos combinados.
# 2FN: Borrar columnas que no son del negocio de dicha tabla, por ejemplo en una tabla alumno tener idColegio con nombreColegio, lo correcto es dejar el id y sacar el nombre.
# 3FN: Eliminar subgrupos de datos en multiples columnas de la tabla, y creando tablas y relaciones nuevas, por ejemplo, si dentro de la tabla Persona tenemos direccion, y 
# dentro de la misma tabla tenemos direccionId, provinciaId, localidadId, etc, esto hace que localidadId teoricamente dependa de la persona, cuando la realidad es que
# tendria que ser un atributo de una nueva tabla direccion con provinciaId y localidadId.

# Tablas temporales
# Son para crear temporalmente tablas desde resultados que se veran repetidos con una finalidad, para evitar que el DBMS repita tantas veces la misma sentencia, para realizar
# pruebas, consultas, analisis, stagear resultados, etc, y/o para evitar que se usen muchos joins en una misma consulta. Se pueden crear como resultados de una query.

# Indices
# Son un mecanismo para acelerar y optimizar las consultas SQL para mejorar los tiempos de respuesta en queries complejas. Mejoran el acceso a datos al darnos una ruta
# mas directa a los registros.

use movies_db;
select * from genres;

# Punto 1
insert into movies (created_at, updated_at, title, rating, awards, release_date, length, genre_id)
values ("2021-03-12","2021-12-15","Spiderman No Way Home", 5, 2, "2021-12-16", 150, 8);

# Punto 2
insert into genres (created_at, updated_at, name, ranking, active)
values ("2021-03-12","2021-12-15","Fantastic Adventures", 13, 1);

# Punto 3
select g.id from genres g where g.name like "%Fantastic Adventures%";
select m.id from movies m where m.title like "%Spiderman No Way Home%";
update movies
set genre_id = 15
where id = 22;

# Punto 4
select * from actors;
# Ya que tiene seteada en null la pelicula favorita, se la vamos a agregar a Bryan Cranston.
update actors
set favorite_movie_id = 22
where id = 47;

# Punto 5
create temporary table movies_copy_bis(
id int not null, 
created_at datetime, 
updated_at datetime, 
title varchar(100) not null, 
rating decimal(2,1) not null, 
awards int not null, 
release_date datetime not null, 
length int, 
genre_id int)
select * from movies;

# Punto 6
select * from movies_copy_bis;
set SQL_SAFE_UPDATES = 0;
delete from movies_copy_bis where awards < 5;

# Punto 7
select g.name from genres g inner join movies m on g.id = m.genre_id group by g.id having count(g.id) > 1;

# Punto 8
select a.first_name, a.last_name from actors a inner join movies m on a.favorite_movie_id = m.id where m.awards > 3;

# Punto 9
explain select * from movies_copy_bis;
explain delete from movies_copy_bis where awards < 7;

# Punto 10
create index actors_idx on actors (id);
create index movie_name on movies (title);

# Punto 11
show index from actors;
show index from movies;
