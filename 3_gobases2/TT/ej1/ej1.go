package ej1

import (
	"fmt"
	"time"
)

type Student struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    time.Time
}

func Ej1() {
	student := Student{"Leandro", "Nicolau", 40684995, time.Now().AddDate(1997, 10, 3)}
	fmt.Println(student)
	fmt.Println(student.Details())
}

func (student Student) Details() string {
	return fmt.Sprintf("\nNombre: %s\nApellido: %s\nDNI: %v\nFecha: %s", student.Nombre, student.Apellido, student.DNI, student.Fecha)
}
