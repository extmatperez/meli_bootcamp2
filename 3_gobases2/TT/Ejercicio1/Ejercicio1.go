package main

import "fmt"

type Estudiante struct {
	Nombre        string `json:"nombre_estudiante`
	Apellido      string `json:"apellido_estudiante"`
	DNI           string `json:"dni_estudiante"`
	Fecha_Ingreso Fecha  `json:"fecha_ingreso_estudiante"`
}

type Fecha struct {
	Dia int `json:"dia"`
	Mes int `json:"mes"`
	Ano int `json:"ano"`
}

func (e Estudiante) getDetails() {
	fmt.Printf("Nombre: \t%s\nApellido: \t%s\nDNI: \t%s\nFecha de ingreso: \t%d/%d/%d\n", e.Nombre, e.Apellido, e.DNI, e.Fecha_Ingreso.Dia, e.Fecha_Ingreso.Mes, e.Fecha_Ingreso.Ano)
}

func main() {
	e0 := Estudiante{"Rodrigo", "Vega Gimenez", "39494914", Fecha{27, 1, 1996}}
	e0.getDetails()
}
