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

func animal(animal string) (func() int, error) {
	switch animal {
	case gato:
		return comidaGato, nil
	case perro:
		return comidaPerro, nil
	case hamster:
		return comidaHamster, nil
	case tarantula:
		return comidaTarantula, nil
	default:
		return nil, errors.New("No existe animal")
	}

}

func comidaPerro() int {
	return 10
}

func comidaGato() int {
	return 5
}
func comidaHamster() int {
	return 250
}
func comidaTarantula() int {
	return 150
}

func main() {
	funPerro, err := animal(perro)
	funGato, err := animal(gato)
	funHamster, err := animal(hamster)
	funTarantula, err := animal(tarantula)

	if err != nil {
		fmt.Printf("ERROR: %v \n", err)
	} else {
		comidaPerr := funPerro()
		comidaGa := funGato()
		comidaHam := funHamster()
		comidatar := funTarantula()

		fmt.Printf("Perro: %v\n", comidaPerr)
		fmt.Printf("Gato: %v\n", comidaGa)
		fmt.Printf("Hamster: %v\n", comidaHam)
		fmt.Printf("Tarantula: %v\n", comidatar)

	}

}
