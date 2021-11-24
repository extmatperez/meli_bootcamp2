package main

import "fmt"

// Ejercicio 5 - Listado nombres

// Una profesora de la universidad quiere tener un listado con todos sus estudiantes. Es necesario crear una aplicación que contenga dicha lista.
// Estudiantes:
// Benjamin, Nahuel, Brenda, Marcos, Pedro, Axel, Alez, Dolores, Federico, Hernán, Leandro, Eduardo, Duvraschka.
 
// Luego de 2 clases, se sumó un estudiante nuevo. Es necesario agregarlo al listado, sin modificar el código que escribiste inicialmente.
// Estudiante:
// Gabriela

func main () {

	var students = []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}

	fmt.Printf("List of students: %s\n", students)

	new_student_name := "Gabriela"

	students = append(students, new_student_name)

	fmt.Printf("New list of students: %s", students)


}