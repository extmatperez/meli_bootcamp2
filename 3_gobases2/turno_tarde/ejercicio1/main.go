// Ejercicio 1 - Registro de estudiantes
// Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

// Nombre: [Nombre del alumno]
// Apellido: [Apellido del alumno]
// DNI: [DNI del alumno]
// Fecha: [Fecha ingreso alumno]

// Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
// Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

package main

import (
	"fmt"
	"reflect"
)

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (d Alumno) detail() {

	//El typeOf de reflect me sirve para utilizar otros metodos como NumField que sirve para saber cuantos campos
	// o Field para acceder al campo
	tipo := reflect.TypeOf(d)
	valor := reflect.ValueOf(d)
	for i := 0; i < tipo.NumField(); i++ {
		fmt.Printf("%v: %v\n", tipo.Field(i).Name, valor.Field(i))
	}
}

func main() {
	student1 := Alumno{
		Nombre:   "Franco",
		Apellido: "Andrada",
		DNI:      41475408,
		Fecha:    "27/08/1998",
	}

	student1.detail()
}
