package main

import (
	"fmt"
)

type estudiante struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func main() {
	e1 := estudiante{"jose", "perez", 123456, "25/22/2021"}
	fmt.Println(e1)
}
