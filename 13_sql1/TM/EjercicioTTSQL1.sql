use movies_db;

# Este es el ejercicio de la tarde de SQL I.

# Punto 1
select * from movies;

# Punto 2
select first_name, last_name, rating from actors;

# Punto 3
select s.title as 'Titulo' from series s;

# Punto 4
select first_name, last_name from actors where rating > 7.5;

# Punto 5
select title, rating, awards from movies where rating > 7.5 and awards > 2;

# Punto 6
select title, rating from movies order by rating asc;

# Punto 7
select title from movies limit 3;

# Punto 8
select * from movies order by rating desc limit 5;

# Punto 9
select * from movies order by rating desc limit 5, 5;

# Punto 10
select * from actors limit 0, 10;

# Punto 11
select * from actors limit 20, 10;

# Punto 12
select * from actors limit 40, 10;

# Punto 13
select title, rating from movies where title like "%Toy Story%";

# Punto 14
select * from actors where first_name like "Sam%";

# Punto 15
select title from movies where year(release_date) between 2004 and 2008;

# Punto 16
select title from movies where rating > 3 and awards > 1 and year(release_date) between 1988 and 2009 order by rating desc;

# Punto 17
select title from movies where rating > 3 and awards > 1 and year(release_date) between 1988 and 2009 order by rating desc limit 10, 3;