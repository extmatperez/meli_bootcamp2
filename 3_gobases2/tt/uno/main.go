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

func (s Student) details() {
	fmt.Println("Name: ", s.Name)
	fmt.Println("Lastname: ", s.LName)
	fmt.Println("DNI : ", s.Dni)
	fmt.Println("Birth Day: ", s.BDay)
}

func main() {
	students := []Student{}
	for _, student := range students {
		student.details()
	}
}
