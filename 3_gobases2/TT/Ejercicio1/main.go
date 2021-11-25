package main

import (
	"fmt"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) mostrarEstudiantes() {
	fmt.Println("Nombre: ", a.Nombre)
	fmt.Println("Apellido: ", a.Apellido)
	fmt.Println("DNI: ", a.DNI)
	fmt.Println("Fecha: ", a.Fecha)
	fmt.Println("*******************************")
}

func main() {
	a1 := Alumno{"Nicolas", "Aponte", 1234567, "10-05-1998"}
	a2 := Alumno{"Maria", "Ardila", 32432, "15-05-1998"}

	a1.mostrarEstudiantes()
	a2.mostrarEstudiantes()

}
