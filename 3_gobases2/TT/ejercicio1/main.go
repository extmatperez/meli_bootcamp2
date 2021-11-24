/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir
el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha
y que tenga un método detalle
*/

package main

import (
	"fmt"
	"strconv"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (a Alumno) detalle() string {
	return ("Nombre: " + a.Nombre + ", Apellido: " + a.Apellido + ", DNI: " + strconv.Itoa(a.DNI) + ", Fecha: " + a.Fecha)
}

func main() {
	a := Alumno{
		Nombre:   "Benjamin",
		Apellido: "Conti",
		DNI:      39620577,
		Fecha:    "15/11/2021",
	}

	fmt.Println(a.detalle())
}
