package main

import (
	"fmt"
)

// Ejercicio 1 - Registro de estudiantes
// Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:
// Nombre: [Nombre del alumno]Apellido: [Apellido del alumno]DNI: [DNI del alumno]Fecha: [Fecha ingreso alumno]

// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
// Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

type Student struct {
	Name    string
	Surname string
	DNI     int
	Date    string
}

func (s *Student) details() {
	fmt.Printf("Student info: \n Name: %s %s\n DNI: %d\n Date: %s\n", s.Name, s.Surname, s.DNI, s.Date)

}

func main() {

	student1 := Student{Name: "John", Surname: "Foo", DNI: 34567890, Date: "March 3 / 2021"}

	student1.details()

}
