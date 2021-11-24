package main

import "fmt"

type fecha struct {
	Dia int
	Mes int
	Año int
}

type Alumno struct{
	Nombre string
	Apellido string
	DNI int
	Fecha fecha
}

func (a Alumno) detalle(){
	fmt.Printf("\nNombre: %v", a.Nombre)
	fmt.Printf("\nApellido: %v", a.Apellido)
	fmt.Printf("\nDNI: %v", a.DNI)
	fmt.Printf("\nFecha: %v", a.Fecha)
}

func main(){

	a1 := Alumno{
		Nombre: "Nahuel", 
		Apellido: "Quinteros",
		DNI: 41525666,
		Fecha: fecha{15,05,1999},
	}

fmt.Println("Datos del alumno:")
a1.detalle()

}