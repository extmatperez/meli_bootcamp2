package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) presentar() {
	fmt.Println("Hola, me llamo", a.Nombre, a.Apellido, "y mi cumple es en", a.Fecha)
}

func main() {

	fmt.Println("Bienvenidos al ejercicio 1")

	alumno1 := Alumno{"Patricio", "Pallua", 1234, "12/12/1997"}
	fmt.Println("El alumno1", alumno1)

	alumno2 := Alumno{
		Nombre:   "Pedro",
		Apellido: "Pallua",
		DNI:      4321,
		Fecha:    "07/04/2003",
	}
	fmt.Println("El alumno2", alumno2)

	alumno2.presentar()
}
