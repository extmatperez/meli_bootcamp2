// Ejercicio 5 - Listado de nombres

// a Una profesora de la universidad quiere tener un listado con todos sus estudiantes. Es necesario crear una aplicación que contenga dicha lista.
// Estudiantes:
// Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.

// b Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, sin modificar el código que escribiste inicialmente.
// Estudiante:
// Gabriela

package main

import "fmt"

func main() {
	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Printf("\nStudents: %v \n", students)
	fmt.Println("\nLuego de 2 clases...")
	new_student := "Gabriela"
	students = append(students, new_student)

	fmt.Printf("\nStudents: %v\n", students)
}
