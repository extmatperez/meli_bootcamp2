use movies_db;

select * from actors;

select se.title, gen.name from series as se inner join genres as gen on se.genre_id = gen.id;

select ep.title, ac.first_name, ac.last_name from episodes as ep inner join actor_episode as ae inner join actors as ac on ae.actor_id = ac.id; 


select se.title, count(sea.number) from series as se inner join seasons as sea on se.id = sea.serie_id group by se.title;


select gen.name, count(mo.title) from genres as gen inner join movies as mo on gen.id = mo.genre_id group by gen.name having count(mo.title) > 3;

select distinct ac.first_name, ac.last_name from actors as ac inner join movies as mo on mo.title like '%La Guerra de las galaxias%';

