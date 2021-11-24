package main

import (
	"fmt"
	"time"
)

func main() {
	alumnos := CargarAlumnos()
	for _, alumn := range alumnos {
		alumn.Detalle()
	}
}

type Alumno struct {
	Nombre       string
	Apellido     string
	DNI          string
	FechaIngreso time.Time
}

func CargarAlumnos() []Alumno {
	alumno1 := Alumno{"Andres", "Ghione", "46502620", time.Date(2015, 01, 15, 0, 0, 0, 0, time.UTC)}
	alumno2 := Alumno{"Maria", "Topaz", "27589409", time.Date(2017, 10, 19, 0, 0, 0, 0, time.UTC)}
	alumno3 := Alumno{"Mario", "Perez", "37284920", time.Date(2021, 12, 26, 0, 0, 0, 0, time.UTC)}
	alumno4 := Alumno{"Valeria", "Smith", "17367389", time.Date(2010, 04, 05, 0, 0, 0, 0, time.UTC)}
	alumno5 := Alumno{"Juan", "Gomez", "19938569", time.Date(2009, 07, 02, 0, 0, 0, 0, time.UTC)}
	alumno6 := Alumno{"Virginia", "Frank", "45679230", time.Date(2017, 11, 18, 0, 0, 0, 0, time.UTC)}
	alumnosExistentes := []Alumno{alumno1, alumno2, alumno3, alumno4, alumno5, alumno6}
	return alumnosExistentes
}

func (alumno *Alumno) Detalle() {
	fmt.Printf("\nNombre: %v\nApellido: %v\nDNI: %v\nFecha: %v\n\n", alumno.Nombre, alumno.Apellido, alumno.DNI, alumno.FechaIngreso)
}
