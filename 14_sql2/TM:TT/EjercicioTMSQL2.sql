# Punto 1
select s.title, g.name from series s inner join genres g on s.genre_id = g.id;

# Punto 2
select e.title, a.first_name, a.last_name from episodes e inner join actor_episode ae on e.id = ae.episode_id
inner join actors a on ae.actor_id = a.id;

# Punto 3
select s.title, count(se.id) from series s inner join seasons se on s.id = se.serie_id group by s.id;

# Punto 4
select g.name, count(m.id) as quantity from genres g inner join movies m on g.id = m.genre_id group by g.id having quantity >= 3;

# Punto 5
select distinct a.first_name, a.last_name from actors a join actor_movie am on a.id = am.actor_id
join movies m on am.movie_id = m.id where m.id in (select mo.id from movies mo where mo.title like "%Guerra de las galaxias%");
