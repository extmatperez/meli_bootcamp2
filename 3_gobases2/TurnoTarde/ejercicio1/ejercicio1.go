/*Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos ç
de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y
que tenga un método detalle*/

package main

import "fmt"

type Alumnos struct {
	nombre   string
	apellido string
	DNI      int
	fecha    string
}

func (a *Alumnos) detalle() {
	fmt.Printf("Alumno %s %s\nDNI: %d\nFecha de Ingreso: %s\n", a.nombre, a.apellido, a.DNI, a.fecha)
}

func main() {
	a1 := Alumnos{
		nombre:   "Alberto",
		apellido: "Gonzalez",
		DNI:      35678453,
		fecha:    "12/11/2020",
	}

	a1.detalle()
}
