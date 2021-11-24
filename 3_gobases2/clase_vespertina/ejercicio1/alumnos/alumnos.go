package alumnos

import "fmt"

type Alumno struct {
	Nombre       string
	Apellido     string
	DNI          string
	FechaIngreso string
}

func CargarAlumnos() []Alumno {
	alumno1 := Alumno{"Andres", "Ghione", "46502620", "15/01/2015"}
	alumno2 := Alumno{"Maria", "Topaz", "27589409", "19/10/2017"}
	alumno3 := Alumno{"Mario", "Perez", "37284920", "26/12/2021"}
	alumno4 := Alumno{"Valeria", "Smith", "17367389", "05/04/2010"}
	alumno5 := Alumno{"Juan", "Gomez", "19938569", "02/07/2009"}
	alumno6 := Alumno{"Virginia", "Frank", "45679230", "18/11/2017"}
	alumnosExistentes := []Alumno{alumno1, alumno2, alumno3, alumno4, alumno5, alumno6}
	return alumnosExistentes
}

func (alumno *Alumno) Detalle() {
	fmt.Printf("Nombre: %v\nApellido: %v\nDNI: %v\nFecha: %v\n", alumno.Nombre, alumno.Apellido, alumno.DNI, alumno.FechaIngreso)
}
