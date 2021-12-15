
-- 2
select * from movies;

-- 3
select first_name, last_name, rating from actors;

-- 4

-- 5
select first_name, last_name from actors 
where rating > 7.5;

-- 6
-- Mostrar el título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios
select title, rating, awards from movies 
where rating > 7.5 and awards > 2;

-- 7
-- Mostrar el título de las películas y el rating ordenadas por rating en forma ascendente.
select title, rating from movies 
order by rating asc;

-- 8
select title from movies
limit 3;

-- 9 mejores 5 pelis segun rating 
select title, rating from movies
order by rating desc limit 5;

-- 10 de la 5 a la 10 con mayor rating
select title, rating from movies
order by rating desc 
limit 5 offset 6;


-- 11
select * from actors limit 10;

-- 12
select first_name, last_name from actors limit 10 offset 30;

-- 13
select first_name, last_name from actors limit 10 offset 50;


-- 14
select title, rating from movies
where title like "%Toy Story%";

-- 15
select * from actors
where first_name like "Sam%";

-- 16
select title, release_date from movies
where release_date between DATE("2004-01-01") and DATE("2008-12-31"); 

-- 17
-- Traer el título de las películas con el rating mayor a 3, con más de 1 premio y
-- con fecha de lanzamiento entre el año 1988 al 2009. Ordenar los resultados por rating.
select title from movies
where rating > 3 and awards > 1 
and release_date between DATE("1988-01-01") and DATE("2009-12-31")
order by rating desc;

-- 18 
-- el top 3 de la consulta anterior despues del 10cimo registro
select title, rating from movies
where rating > 3 and awards > 1 
and release_date between DATE("1988-01-01") and DATE("2009-12-31")
order by rating desc 
limit 3 offset 10;


