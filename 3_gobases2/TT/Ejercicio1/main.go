/* Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de
cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método
detalle
*/

package main

import "fmt"

type Alumnos struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (v *Alumnos) detalle() {
	fmt.Printf("\nNombre: %v\nApellido: %v\nDNI: %v\nFecha: %v\n", v.Nombre, v.Apellido, v.DNI, v.Fecha)
}

func main() {
	p1 := Alumnos{"Julian", "Alvarez", 38457216, "13/03/1996"}
	p2 := Alumnos{"Wanchope", "Avila", 37412325, "14/10/1989"}
	p1.detalle()
	p2.detalle()
}
