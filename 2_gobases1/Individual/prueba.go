package main

import (
	"fmt"
	"reflect"
)

type Persona struct {
	Nombre   string //`json: "nombrecito"`
	Apellido string `json: "apellidito"`
	Edad     int    `json: "edad"`
}

func main() {
	p1 := Persona{}
	p1.Nombre = "Juan"
	persona := Persona{}
	p := reflect.TypeOf(persona)
	persona.Nombre = "Juan"

	fmt.Println(p.Field(0))
	fmt.Println(p.Name())
	fmt.Println(p.Kind())
	fmt.Println(p1.Nombre)

}
