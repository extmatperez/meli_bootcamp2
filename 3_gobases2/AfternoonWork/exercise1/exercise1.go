/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

*/
package main

import "fmt"

type Student struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (s Student) detalle() string {
	return fmt.Sprintf("Nombre: %s \nApellido: %s\nDNI: %s\nFecha: %s", s.Nombre, s.Apellido, s.DNI, s.Fecha)
}

func main() {
	s1 := Student{Nombre: "Jose", Apellido: "Rios", DNI: "184965984", Fecha: "24-11-2021"}
	fmt.Println(s1.detalle())
}
