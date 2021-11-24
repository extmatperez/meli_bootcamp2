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

	animalPerro, err := Animal(perro)
	if err != nil {
		fmt.Println("Error con perro")
		return
	}
	animalGato, err := Animal(gato)

	if err != nil {
		fmt.Println("Error con gato")
		return
	}

	animalHamster, err := Animal(hamster)

	if err != nil {
		fmt.Println("Error con hamster")
		return
	}

	animalTarantula, err := Animal(tarantula)

	if err != nil {
		fmt.Println("Error con tarantula")
		return
	}

	var cantidadEnKg float64
	cantidadEnKg += animalPerro(5)
	cantidadEnKg += animalGato(8)
	cantidadEnKg += animalHamster(1)
	cantidadEnKg += animalTarantula(1)

	fmt.Printf("Cantidad necesaria de alimento: %.2f\n", cantidadEnKg)
}

func Animal(tipoAnimal string) (func(c int) float64, error) {

	switch tipoAnimal {
	case perro:
		return perroF, nil
	case gato:
		return gatoF, nil
	case hamster:
		return hamsterF, nil
	case tarantula:
		return tarantulaF, nil
	default:
		return nil, errors.New("No se reconoce ese tipo de animal")
	}
}

func perroF(c int) float64 {
	return float64(c) * 10.0
}

func gatoF(c int) float64 {
	return float64(c) * 5.0
}

func hamsterF(c int) float64 {
	return float64(c) * 0.25
}

func tarantulaF(c int) float64 {
	return float64(c) * 0.15
}
