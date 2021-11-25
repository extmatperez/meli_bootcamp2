package main

import (
	"fmt"
)

type Fecha struct {
	Dia  int
	Mes  int
	Anio int
}
type Estudiante struct {
	Nombre       string
	Apellido     string
	Dni          string `json: "DNI"`
	FechaIngreso Fecha  `json: "Fecha"`
}

func main() {

	alumno1 := Estudiante{
		Nombre:   "Carlos",
		Apellido: "Perez",
		Dni:      "40404040",
		FechaIngreso: Fecha{
			Dia:  24,
			Mes:  11,
			Anio: 2020,
		},
	}
	alumno2 := Estudiante{
		Nombre:   "Marta",
		Apellido: "Juarez",
		Dni:      "302020302",
		FechaIngreso: Fecha{
			Dia:  21,
			Mes:  1,
			Anio: 2021,
		},
	}

	details(alumno1)
	details(alumno2)

}

func details(e Estudiante) {
	fmt.Printf("Nombre: [%s]\n", e.Nombre)
	fmt.Printf("Apellido: [%s]\n", e.Apellido)
	fmt.Printf("DNI: [%s] \n", e.Dni)
	fmt.Println("Fecha de Ingreso:", e.FechaIngreso.Dia, "/", e.FechaIngreso.Mes, "/", e.FechaIngreso.Anio)
}
