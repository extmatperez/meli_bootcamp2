package main

import "fmt"

func main() {

	// a.
	students := []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alex", "Dolores", "Federico", "Hernán", "Leandro", "Eduardo", "Duvraschka"}
	fmt.Println(students)

	// b.
	actualStudents := make([]string, 16)

	actualStudents = students
	actualStudents = append(actualStudents, "Gabriela")

	fmt.Println(actualStudents)

}
