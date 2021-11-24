package main

import "fmt"

func main() {
	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println("Cantidad de estudiantes: ", len(students), "\nEstudiantes actuales: ", students)

	students = append(students, "Gabriela")
	fmt.Println("Cantidad de estudiantes: ", len(students), "\nEstudiantes actuales: ", students)
}
