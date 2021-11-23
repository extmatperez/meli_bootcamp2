package ej5

import "fmt"

var students = []string{
	"Benjamin",
	"Nahuel",
	"Brenda",
	"Marcos",
	"Pedro",
	"Axel",
	"Alez",
	"Dolores",
	"Federico",
	"Hernan",
	"Leandro",
	"Eduardo",
	"Duvraschka",
}

func Ej5() []string {
	fmt.Println(students)
	return students
}

func AddStudent(student string) []string {
	students = append(students, student)
	fmt.Println(students)
	return students
}
