package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	Dni      string
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Println("Nombre:\t\t", a.Nombre)
	fmt.Println("Apellido:\t", a.Apellido)
	fmt.Println("DNI:\t\t", a.Dni)
	fmt.Println("Fecha:\t\t", a.Fecha)
}
