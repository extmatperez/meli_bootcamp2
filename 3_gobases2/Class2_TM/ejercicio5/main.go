package main

import (
	"errors"
	"fmt"
)

/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que
puedan haber muchos más animales que refugiar.

Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el
animal especificado y que retorne una función y un error (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad
del tipo de animal especificado.

*/

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	animalPerro, err := Animal(perro)
	animalGato, err := Animal(gato)
	animalHamster, err := Animal(hamster)
	animalTarantula, err := Animal(tarantula)

	if err != nil {
		fmt.Printf("Hubo un error: %v", err)
	} else {
		var cantidad float64
		cantidad += animalPerro(4)
		cantidad += animalGato(12)
		cantidad += animalHamster(3)
		cantidad += animalTarantula(6)

		fmt.Printf("Necesita %v kg de alimento \n", cantidad)
	}
}

func Animal(valor string) (func(cant int) float64, error) {
	switch valor {
	case perro:
		return animalPerro, nil
	case gato:
		return animalGato, nil
	case hamster:
		return animalHamster, nil
	case tarantula:
		return animalTarantula, nil
	default:
		return nil, errors.New("Invalid Operator")
	}
}

func animalPerro(cant int) float64 {
	alimento := 10.0 * float64(cant)
	return alimento
}

func animalGato(cant int) float64 {
	alimento := 5.0 * float64(cant)
	return alimento
}

func animalHamster(cant int) float64 {
	alimento := 0.250 * float64(cant)
	return alimento
}

func animalTarantula(cant int) float64 {
	alimento := 0.150 * float64(cant)
	return alimento
}
