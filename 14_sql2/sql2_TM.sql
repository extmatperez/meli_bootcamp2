#1. ¿A qué se denomina JOIN en una base de datos?
	/* el JOIN es una sentencia SQL que permite combinar registros entre multiples tabla  */
    
#2. ¿A qué se denomina JOIN en una base de datos?
	/* 
		el INNER JOIN combina y trae los registros que tengan datos comunes entre dos tabla  mientras que un 
		LEFT JOIN trae todos los datos de la primera table y si tiene datos la segunda los trae, si no tienen
        coincidencia trae el campo como null
	*/
    
#3. ¿Para qué se utiliza el GROUP BY?
	/*agrupar los datos de las consultas por campos determinado luego de la sentencia */

#4. ¿Para qué se utiliza el HAVING?
	/*
		HAVING tiene la misma utilidad que el WHERE con la diferencia que HAVING se usa para datos que fueron 
		agrupados mientras que el WHERE se usa para datos comunes sin agrupamiento
	*/
    
#5. Dado lo siguientes diagramas indique a qué tipo de JOIN corresponde cada uno:
	/* INNER JOIN y LEFT JOIN*/

#6 Escriba una consulta genérica por cada uno de los diagramas a continuación:

	/*
		SELECT * 
        FROM `tabla_1`
		RIGHT JOIN `tabla_2` ON `tabla_1.id_tabla2` = `tabla_2.id`
        
        SELECT * 
        FROM `tabla_1`
		FULL JOIN `tabla_2` ON `tabla_1.id_tabla2` = `tabla_2.id`
    */