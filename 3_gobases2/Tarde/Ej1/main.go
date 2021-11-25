/*
? Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para
imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los
alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido,
DNI, Fecha y que tenga un método detalle

*/

package main

import "fmt"

func main() {

	type Alumno struct {
		Nombre   string
		Apellido string
		Edad     int
		Fecha    string
	}

	alumno1 := Alumno{"Nico", "Arguello", 36, "08/08/85"}

	alumno2 := Alumno{
		Nombre:   "Cele",
		Apellido: "Gonzalez",
		Edad:     24,
		Fecha:    "22/01/98",
	}

	fmt.Printf("%+v\n", alumno1)
	fmt.Printf("%+v", alumno2)

	//func (v Alumno) detalle() {

	//}
}
