package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Printf("\nNombre: %v, Apellido: %v, DNI: %v, Fecha: %v \n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {

	alumno1 := Alumno{"Carlos", "Davila", "789651", "12/11/2021"}
	alumno1.detalle()

	alumno2 := Alumno{
		Nombre:   "Pedro",
		Apellido: "Perez",
		DNI:      "1234567",
	}
	alumno2.detalle()

	var alumno3 Alumno
	alumno3.Nombre = "Andres"
	alumno3.Apellido = "Gonzales"
	alumno3.DNI = "976543"
	alumno3.Fecha = "22/11/2021"
	alumno3.detalle()

}
