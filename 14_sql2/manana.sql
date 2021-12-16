use movies_db;

select se.title, ge.name from  series se join genres ge on se.genre_id = ge.id ;

select ac.first_name,ac.last_name,ep.title from actor_episode ae join actors ac on ae.actor_id = ac.id join episodes ep on ep.id= ae.episode_id;

select se.title, max(sea.number) from series se join seasons sea on sea.serie_id = se.id group by title;

select   ge.name, count(ge.name) as suma from genres ge join movies mo on mo.genre_id = ge.id  group by ge.name having suma >=3 ;

select ac.first_name, ac.last_name from actors ac join actor_movie am on ac.id = am.actor_id join movies mv on mv.id = am.movie_id  where mv.title like 'La Guerra de las galaxia%'  group by ac.id ;