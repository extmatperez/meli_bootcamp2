package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	animalPerro, err := animal(perro)
	animalGato, err := animal(gato)
	animalHamster, err := animal(hamster)
	animalTarantula, err := animal(tarantula)

	if err != nil {
		fmt.Printf("Hubo un error: %v \n", err)
	} else {
		var cantidad, cantidad1, cantidad2, cantidad3, cantidad4 float64
		cantidad += animalPerro(5)
		cantidad1 += animalGato(5)
		cantidad2 += animalHamster(10)
		cantidad3 += animalTarantula(10)
		cantidad4 = cantidad + cantidad1 + cantidad2 + cantidad3

		fmt.Printf("Necesita %v kg de alimento para perros\n", cantidad)
		fmt.Printf("Necesita %v kg de alimento para gatos\n", cantidad1)
		fmt.Printf("Necesita %v kg de alimento para hamsters\n", cantidad2)
		fmt.Printf("Necesita %v kg de alimento para tarantulas\n", cantidad3)

		fmt.Printf("Necesita %v kg de alimento en total \n", cantidad4)

	}

}

func animal(animal string) (func(q int) float64, error) {
	switch animal {
	case perro:
		return perroFunc, nil
	case gato:
		return gatoFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return taranFunc, nil
	default:
		return nil, errors.New("operacion invalida")

	}

}

func perroFunc(q int) float64 {
	var comida = float64(q) * 10.0
	return comida
}

func gatoFunc(q int) float64 {
	var comida = float64(q) * 5.0
	return comida
}

func hamsterFunc(q int) float64 {
	var comida = float64(q) * 0.250
	return comida
}

func taranFunc(q int) float64 {
	var comida = float64(q) * 0.150
	return comida
}
