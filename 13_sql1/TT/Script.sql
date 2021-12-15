select * from movies 

select first_name, last_name,rating from actors 

select title as titulo from series 

select first_name as Nombre, last_name as Apellido from actors where rating > 7.5

select m.title as Titulo, m.rating as Ranking, m.awards as Premios from movies m where m.rating > 7.5 and m.awards > 2

select m.title as Titulo, m.rating as Ranking from movies m order by m.rating ASC  

select m.title as Titulo from movies m limit 3

select * from movies m order by m.rating DESC limit 5 

select * from movies m order by m.rating DESC limit 5 offset 5

select * from actors a limit 10

select * from actors a limit 10 offset 30

select * from actors a limit 10 offser 50

select m.title as Titulo, m.rating as Ranking from movies m  where m.title LIKE "Toy Story%"

select * from actors a WHERE a.first_name LIKE 'Sam%' 

select m.title as Titulo from movies m where m.release_date BETWEEN '2004-01-01 00:00:00' and '2008-01-01 00:00:00';

select m.title as Titulo from movies m where m.rating > 3 and m.awards > 1 and m.release_date BETWEEN '1988-01-01 00:00:00' and '2009-01-01 00:00:00' ORDER BY m.rating DESC ;

select m.title as Titulo from movies m where 
m.rating > 3 
and m.awards > 1
and m.release_date 
BETWEEN '1988-01-01 00:00:00' and '2009-01-01 00:00:00' 
ORDER BY m.rating DESC
LIMIT 3