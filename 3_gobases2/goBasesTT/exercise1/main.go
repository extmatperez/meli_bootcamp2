package main

import (
	"fmt"
	"time"
)

func main() {
	var nombre, apellido string
	var dni int

	fmt.Println("Ejercicio 1")
	fmt.Println("Ingrese nombre")
	fmt.Scanf("%s", &nombre)
	fmt.Println("Ingrese apellido")
	fmt.Scanf("%s", &apellido)
	fmt.Println("Ingrese DNI")
	fmt.Scanf("%v", &dni)

	alumno := Alumno{nombre, apellido, dni, time.Now()}
	fmt.Println(alumno)
	fmt.Println(alumno.detalles())

}

type Alumno struct {
	Nombre, Apellido string
	Dni              int
	Fecha            time.Time
}

func (alumno Alumno) detalles() string {
	return fmt.Sprintf("\nNombre: %s\nApellido: %s\nDNI: %v\nFecha: %s", alumno.Nombre, alumno.Apellido, alumno.Dni, alumno.Fecha)
}
