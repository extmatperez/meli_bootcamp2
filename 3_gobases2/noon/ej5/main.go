package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	tarantula = "tarantula"
	hamster   = "hamster"
)

func animalPerro(cantidad int) float64 {
	return float64(cantidad * 10)
}

func animalGato(cantidad int) float64 {
	return float64(cantidad * 5)
}

func animalTarantula(cantidad int) float64 {
	return float64(cantidad) * 0.15
}

func animalHamster(cantidad int) float64 {
	return float64(cantidad) * 0.25
}

func animal(nombre string) (func(cantidad int) float64, error) {
	switch nombre {
	case perro:
		return animalPerro, nil
	case gato:
		return animalGato, nil
	case tarantula:
		return animalTarantula, nil
	case hamster:
		return animalHamster, nil
	default:
		return nil, errors.New("el animal no esta en el refugio")
	}
}

func pruebaAnimal(nombre string, cantidad int) {
	animalFunc, err := animal(nombre)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El valor devuelto por la operacion %s es de %0.2f\n", nombre, animalFunc(cantidad))
	}
}

func main() {
	pruebaAnimal(gato, 10)
	pruebaAnimal(perro, 25)
	pruebaAnimal(tarantula, 30)
	pruebaAnimal(hamster, 50)
	pruebaAnimal("Otro animal", 50)
}
