package main

import {
	"golang.org/x/text/date"
	"fmt"
}

func main() {
	type Alumno struct {
		Nombre       string
		Apellido     string
		Dni          string
		FechaIngreso date
	}

	fmt.Println("\nIngrese el nombre del alumno")
	fmt.Scanln(&Nombre)
	fmt.Println("\nIngrese el apellido del alumno")
	fmt.Scanln(&Apellido)
	fmt.Println("\nIngrese el dni del alumno")
	fmt.Scanln(&Dni)
	fmt.Println("\nIngrese la fecha de ingreso")
	fmt.Scanln(&FechaIngreso)

}
