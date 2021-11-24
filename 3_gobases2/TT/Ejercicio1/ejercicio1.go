package main

import "fmt"

func main() {
	var students []Student

	exit := 1
	for exit != 0 {
		var studentName string
		var studentLastname string
		var studentDni int
		var studentDate string
		fmt.Println("Ingrese el nombre del estudiante")
		fmt.Scanf("%s", &studentName)
		fmt.Println("Ingrese el apellido del estudiante")
		fmt.Scanf("%s", &studentLastname)
		fmt.Println("Ingrese el DNI del estudiante")
		fmt.Scanf("%d", &studentDni)
		fmt.Println("Ingrese la fecha del estudiante")
		fmt.Scanf("%s", &studentDate)

		student := Student{
			Name:     studentName,
			Lastname: studentLastname,
			Dni:      studentDni,
			Date:     studentDate,
		}

		students = append(students, student)

		fmt.Println("")
		fmt.Println("==========Estudiantes actuales==========")
		for i := 0; i < len(students); i++ {
			students[i].details()
			fmt.Println("")
		}
		fmt.Println("========================================")
		fmt.Println("")

		fmt.Println("* Si desea salir, ingrese 0. En caso contrario, presione cualquier otra tecla")
		fmt.Scanf("%d", &exit)
	}
}

type Student struct {
	Name     string
	Lastname string
	Dni      int
	Date     string
}

func (s Student) details() {
	fmt.Println("Nombre:", s.Name)
	fmt.Println("Apellido:", s.Lastname)
	fmt.Println("DNI:", s.Dni)
	fmt.Println("Fecha:", s.Date)
}
