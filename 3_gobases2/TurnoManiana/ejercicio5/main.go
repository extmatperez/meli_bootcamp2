package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantura"
)

func animal(tipoAnimal string) (float64, error) {
	var cantidad float64
	switch tipoAnimal {
	case perro:
		cantidad = 10
		return funcPerro(cantidad), nil
	case gato:
		cantidad = 5
		return funcGato(cantidad), nil
	case hamster:
		cantidad = 15
		return funcHamster(cantidad), nil
	case tarantula:
		cantidad = 3
		return funcTarantula(cantidad), nil
	default:
		return 0, errors.New("Animal no existe")
	}
}

func funcPerro(cantidadAnimales float64) float64 {
	cantAlimento := 10.00
	return cantAlimento * cantidadAnimales
}

func funcGato(cantidadAnimales float64) float64 {
	cantAlimento := 5.00
	return cantAlimento * cantidadAnimales
}

func funcHamster(cantidadAnimales float64) float64 {
	cantAlimento := 0.250
	return cantAlimento * cantidadAnimales
}

func funcTarantula(cantidadAnimales float64) float64 {
	cantAlimento := 0.150
	return cantAlimento * cantidadAnimales
}

func main() {

	resultado, err := animal(tarantula)
	if err != nil {
		fmt.Println("No existe el animal")
	} else {
		fmt.Printf("El animal consumira %vKG. de alimento\n", resultado)
	}
}
