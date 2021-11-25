package main

import (
	"fmt"
	"strconv"
)

type Date struct {
	year  int
	month int
	day   int
}

type Alumnos struct {
	Name     string
	LastName string
	DNI      int
	admDate  Date
}

func (alumno Alumnos) detalle() string {
	return "\n" + "Nombre: " + alumno.Name + "\n" + "Apellido: " + alumno.LastName + "\n" + "DNI: " + strconv.Itoa(alumno.DNI) + "\n" + "Fecha de nacimiento: " + strconv.Itoa(alumno.admDate.day) + "/" + strconv.Itoa(alumno.admDate.month) + "/" + strconv.Itoa(alumno.admDate.year) + "\n\n----------"
}
func main() {
	/*
		Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

		Nombre: [Nombre del alumno]
		Apellido: [Apellido del alumno]
		DNI: [DNI del alumno]
		Fecha: [Fecha ingreso alumno]

		Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
		Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
	*/

	flagReg := "si"
	var student Alumnos
	var students []Alumnos

	// CODIGO PARA CARGAR ALUMNOS

	for flagReg == "si" {
		fmt.Printf("Ingrese el nombre del alumno: ")
		fmt.Scan(&student.Name)

		fmt.Printf("Ingrese el apellido del alumno: ")
		fmt.Scan(&student.LastName)

		fmt.Printf("Ingrese el DNI del alumno: ")
		fmt.Scan(&student.DNI)

		fmt.Printf("Ingrese el dia de nacimiento del alumno: ")
		fmt.Scan(&student.admDate.day)

		fmt.Printf("Ingrese el mes de nacimiento del alumno: ")
		fmt.Scan(&student.admDate.month)

		fmt.Printf("Ingrese el año de nacimiento del alumno: ")
		fmt.Scan(&student.admDate.year)

		students = append(students, student)

		fmt.Println("Desea cargar otro alumno? si/no: ")
		fmt.Scan(&flagReg)
	}

	// CODIGO PARA IMPRIMIR DATOS DE ALUMNOS

	for _, student := range students {
		fmt.Println(student.detalle())
	}
}
