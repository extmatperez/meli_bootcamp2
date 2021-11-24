package main

import "fmt"

type Date struct {
	Day   int
	Month int
	Year  int
}

type Student struct {
	Name  string
	LName string
	Dni   int
	BDay  Date
}

var students []Student

func (s Student) details() {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Lastname: ", s.LName)
	fmt.Println("DNI : ", s.Dni)
	fmt.Println("Birth Day: ", s.BDay)
}

func printStudents(students []Student) {
	for _, student := range students {
		student.details()
		fmt.Println("____________________")
	}
}

func main() {

	students = append(students, Student{"Mario", "Santos", 23389375, Date{18, 1, 1964}})
	students = append(students, Student{"Pablo", "Lamponne", 25639332, Date{28, 10, 1965}})
	students = append(students, Student{"Emilio", "Ravenna", 22287190, Date{3, 12, 1963}})
	students = append(students, Student{"Gabriel", "Medina", 24412610, Date{5, 7, 1964}})

	printStudents(students)
}
