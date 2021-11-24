package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nuevoEstudiante string
	fmt.Println("Ejercicio 5")
	fmt.Println("Escribi el nombre del estudiante que se suma a la clase")
	fmt.Scanf("%s", &nuevoEstudiante)
	fmt.Println(clases(nuevoEstudiante))

}

func clases(nuevoEstudiante string) ([]string, string) {
	estudiantes := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	if reflect.ValueOf(nuevoEstudiante).Kind() == reflect.String {
		estudiantes = append(estudiantes, nuevoEstudiante)
		return estudiantes, "\n esta es la lista de tus estudiantes: "
	} else {
		return nil, " tenes que ingresar el numero del mes en cuestion "
	}

}
