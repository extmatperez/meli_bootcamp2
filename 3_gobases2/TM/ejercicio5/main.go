/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
haber muchos más animales que refugiar.

Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal
especificado y que retorne una función y un error (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del
tipo de animal especificado.
*/

package main

import (
	"errors"
	"fmt"
)

const (
	tarantula = "tarantula"
	hamster  = "hamster"
	perro     = "perro"
	gato      = "gato"
)

func animal(tipoAnimal string) (func(cantidad int) float64, error) {

	switch tipoAnimal {
	case "tarantula":
		return cantidadTarantulas, nil
	case "hamster":
		return cantidadHamsters, nil
	case "perro":
		return cantidadPerros, nil
	case "gato":
		return cantidadGatos, nil
	default:
		return nil, errors.New("animal no existente")
	}
}

func cantidadTarantulas(cantidad int) float64 {
	return float64(cantidad * 150)
}

func cantidadHamsters(cantidad int) float64 {
	return float64(cantidad * 250)
}

func cantidadPerros(cantidad int) float64 {
	return float64(cantidad * 10000)
}

func cantidadGatos(cantidad int) float64 {
	return float64(cantidad * 5000)
}

func main() {
	cantidadTotal := 0.0

	animalTarantula, err1 := animal(tarantula)
	animalHamster, err2 := animal(hamster)
	animalPerro, err3 := animal(perro)
	animalGato, err4 := animal(gato)

	if err1 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("%v\n", animalTarantula(5))
		cantidadTotal += animalTarantula(5)
	}

	if err2 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("%v\n", animalHamster(3))
		cantidadTotal += animalHamster(3)
	}

	if err3 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("%v\n", animalPerro(4))
		cantidadTotal += animalPerro(4)
	}

	if err4 != nil {
		fmt.Printf("Error: %v\n", err1)
	} else {
		fmt.Printf("%v\n", animalGato(2))
		cantidadTotal += animalGato(2)
	}

	fmt.Printf("La cantidad total de alimento es: %.2f gramos\n", cantidadTotal)
}