package main

import "fmt"

type fecha struct {
	Dia int
	Mes int
	Año int
}

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    fecha
}

func (a Alumno) detalle() {
	fmt.Printf("Nombre: %v\n Apellido: %v \n DNI: %v\n Fecha de nacimiento: %v\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	alu1 := crearAlumno("Matias", "De Bonis", 43441682, fecha{8, 6, 2001})
	alu1.detalle()
}

func crearAlumno(nombre string, apellido string, dni int, fecha fecha) Alumno {
	return Alumno{nombre, apellido, dni, fecha}
}
