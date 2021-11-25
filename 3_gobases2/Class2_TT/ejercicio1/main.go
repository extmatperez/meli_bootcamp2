package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a1 := Alumno{"Jose", "Villegas", 43874582, Fecha_{23, 4, 1998}}
	a2 := Alumno{"Pepe", "Dias", 34875445, Fecha_{12, 5, 1988}}
	a3 := Alumno{"Sandra", "Noches", 897349873, Fecha_{31, 6, 1978}}
	a4 := Alumno{"Alfonso", "Wainer", 348973986, Fecha_{26, 7, 2000}}

	array := []Alumno{a1, a2, a3, a4}
	for _, alumno := range array {
		miJSON, err := json.Marshal(alumno)

		fmt.Println(string(miJSON))
		fmt.Println(err)
	}
}

type Alumno struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	DNI      int    `json:"dni"`
	Fecha    Fecha_ `json:"fecha"`
}

type Fecha_ struct {
	Dia  int
	Mes  int
	Anio int
}
