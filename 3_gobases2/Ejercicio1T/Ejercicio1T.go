package main

import (
	"fmt"
	"strconv"
)

type Estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (e Estudiante) getNombre() string {
	return "Nombre: " + e.Nombre
}
func (e Estudiante) getApellido() string {
	return "Apellido: " + e.Apellido
}
func (e Estudiante) getDNI() string {
	return "DNI: " + strconv.Itoa(e.DNI)
}

func (e Estudiante) getFecha() string {
	return "Fecha: " + e.Fecha
}
func (e *Estudiante) setNombre(nombre string) {
	e.Nombre = nombre
}
func main() {
	var estudiantes []Estudiante

	alumno1 := Estudiante{"Walter", "Castillo", 12345, "23-02-2021"}
	alumno2 := Estudiante{
		Nombre: "jesus",
		DNI:    4321,
	}
	var alumno3 Estudiante
	alumno3.Nombre = "harol"
	estudiantes = append(estudiantes, alumno1, alumno2, alumno3)
	estudiantes[0].setNombre("Terwal")
	for i := 0; i < len(estudiantes); i++ {
		fmt.Println("Estudiante", i+1, ":")
		fmt.Println("\t\t", estudiantes[i].getNombre())
		fmt.Println("\t\t", estudiantes[i].getApellido())
		fmt.Println("\t\t", estudiantes[i].getDNI())
		fmt.Println("\t\t", estudiantes[i].getFecha())
	}
}
