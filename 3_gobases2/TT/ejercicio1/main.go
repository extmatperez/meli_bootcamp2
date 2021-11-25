package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	Dni      int
	Fecha    string
}

func (a Alumno) detalle() {
	fmt.Println(a)
}

func main() {
	alumno1 := Alumno{"Damian", "Zamora", 34203601, "04/12/2020"}
	alumno1.detalle()

	var alumno2 Alumno
	alumno2.Nombre = "Daniel"
	fmt.Println(alumno2)

	alumno3 := Alumno{"Juan", "Zamora", 34203601, "04/12/2020"}
	alumno3.detalle()
	alumno4 := Alumno{"Marcos", "Zamora", 34203601, "04/12/2020"}
	alumno4.detalle()

}
