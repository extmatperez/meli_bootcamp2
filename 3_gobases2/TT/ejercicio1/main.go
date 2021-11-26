package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func main() {
	a := Alumno{"Seba", "Chiappa", 1231231, "10/08/2018"}
	a.detalle()
}

func (alumno Alumno) detalle() {
	fmt.Printf("Nombre\t\t%s\n", alumno.Nombre)
	fmt.Printf("Apellido:\t%s\n", alumno.Apellido)
	fmt.Printf("DNI:\t\t%d\n", alumno.DNI)
	fmt.Printf("Fecha:\t\t%s\n", alumno.Fecha)
}
