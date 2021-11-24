package main

import "fmt"

type Alumno struct {
	nombre string
	apellido string
	dni string
	fecha string
}

func (alumno Alumno) details() string {
	return fmt.Sprintf("Nombre: [%s]\nApellido: [%s]\nDNI: [%s]\nFecha: [%s]\n", alumno.nombre, alumno.apellido, alumno.dni, alumno.fecha)
}

func main() {
	alumno := Alumno{
		nombre: "Juan",
		apellido: "Perez",
		dni: "12345678",
		fecha: "12/12/2012",
	}
	fmt.Println(alumno.details());
}