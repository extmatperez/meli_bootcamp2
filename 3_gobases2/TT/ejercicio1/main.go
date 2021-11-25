package main

import "fmt"

type Fecha struct {
	Dia  int
	Mes  int
	Anio int
}

type Alumno struct {
	Nombre       string
	Apellido     string
	DNI          int
	FechaIngreso Fecha
}

func (a *Alumno) detalle() {
	fmt.Printf("Nombre: %v \nApellido: %v \nDNI: %v \nFecha: %v\n", a.Nombre, a.Apellido, a.DNI, a.FechaIngreso)
}

func main() {
	al := Alumno{"Ivan", "Arevalo", 40671767, Fecha{10, 10, 1997}}
	al.detalle()
}
