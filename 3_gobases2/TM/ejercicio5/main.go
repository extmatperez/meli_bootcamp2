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

func Animal(animal string) (func(cant int) float64, error) {
	switch animal {
	case perro:
		return animalPerro, nil
	case gato:
		return animalGato, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, errors.New("invalid animal")
	}
}

func animalPerro(cant int) float64 {
	var alimento float64 = 10.0 * float64(cant)
	return alimento
}

func animalGato(cant int) float64 {
	var alimento float64 = 5.0 * float64(cant)
	return alimento
}

func animalHamster(cant int) float64 {
	var alimento float64 = 0.250 * float64(cant)
	return alimento
}

func animalTarantula(cant int) float64 {
	var alimento float64 = 0.150 * float64(cant)
	return alimento
}

func main() {
	animalPerro, err := Animal(perro)
	animalGato, err := Animal(gato)
	animalHamster, err := Animal(hamster)
	animalTarantula, err := Animal(tarantula)

	if err != nil {
		fmt.Printf("Hubo un error: %v", err)
	} else {
		var cantidad float64
		cantidad += animalPerro(6)
		cantidad += animalGato(8)
		cantidad += animalHamster(5)
		cantidad += animalTarantula(8)

		fmt.Printf("Necesita %v kg de alimento \n", cantidad)
	}

}
