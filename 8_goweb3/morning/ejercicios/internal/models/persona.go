package models

type Persona struct {
	ID       	int    		`json:"id"`
	Nombre   	string 		`json:"nombre"`
	Apellido 	string 		`json:"apellido"`
	Edad     	int    		`json:"edad"`
	Domicilio 	Ciudad 		`json:"domicilio"`
}

type Ciudad struct {
	ID 				int			`json:"id"`
	NombreCiudad 	string 		`json:"nombre_ciudad"`
	NombrePais 		string		`json:"nombre_pais"`
}