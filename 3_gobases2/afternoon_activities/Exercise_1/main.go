/* Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as,
de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

package main

import "fmt"

type Date struct {
	Day   int
	Month string
	Year  int
}

type Student struct {
	Name    string
	Surname string
	DNI     int
	Date    Date
}

func main() {
	student_1 := Student{
		Name:    "John",
		Surname: "Doe",
		DNI:     12345678,
		Date: Date{
			Day:   27,
			Month: "July",
			Year:  2021,
		},
	}

	student_2 := Student{
		Name:    "María",
		Surname: "Doe",
		DNI:     87654321,
		Date: Date{
			Day:   03,
			Month: "December",
			Year:  2021,
		},
	}

	fmt.Printf("Nombre: %s\nSurname: %s\nDNI: %d\nDate: %d/%s/%d\n",
		student_1.Name, student_1.Surname, student_1.DNI, student_1.Date.Day, student_1.Date.Month, student_1.Date.Year)
	fmt.Printf("\n")
	fmt.Printf("Nombre: %s\nSurname: %s\nDNI: %d\nDate: %d/%s/%d\n",
		student_2.Name, student_2.Surname, student_2.DNI, student_2.Date.Day, student_2.Date.Month, student_2.Date.Year)
}
