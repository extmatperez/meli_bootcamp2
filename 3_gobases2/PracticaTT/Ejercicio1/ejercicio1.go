package main

import "fmt"

type Fecha struct {
	dia, mes, anio int
}

type Alumno struct {
	nombre   string
	apellido string
	DNI      int
	fechaIng Fecha
}

/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de
los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y
que tenga un método detalle
*/

func (a Alumno) detalle() {
	fmt.Printf("Nombre: %v\n", a.nombre)
	fmt.Printf("Apellido: %v\n", a.apellido)
	fmt.Printf("DNI: %v\n", a.DNI)
	fmt.Printf("Fecha: %v\n", a.fechaIng)
}

func main() {

	alumno1 := Alumno{"Facundo", "Bouza", 41332191, Fecha{21, 3, 2016}}

	alumno1.detalle()

}
