// Ejercicio 5 - Calcular cantidad de alimento

// Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

// Por perro necesitan 10 kg de alimento
// Por gato 5 kg
// Por cada Hamster 250 gramos.
// Por Tarántula 150 gramos.

// Se solicita:
// Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un error (en caso que no exista el animal)
// Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

// ejemplo:

// const (
// 	perro = "perro"
// 	gato = "gato"
//  )

//  ...

//  animalPerro, err := Animal(perro)
//  animalGato, err := Animal(gato)
//  ...

// var cantidad float64
// cantidad += animalPerro(5)
// cantidad += animalGato(8)

package main

import (
	"fmt"
)

func qty_food_perro(qty float64) float64 {
	food_dog := 10000
	return float64(food_dog) * qty
}
func qty_food_gato(qty float64) float64 {
	food_gato := 5000
	return float64(food_gato) * qty
}
func qty_food_tarantula(qty float64) float64 {
	food_tarantula := 150
	return float64(food_tarantula) * qty
}
func qty_food_hamster(qty float64) float64 {
	food_hamster := 5000
	return float64(food_hamster) * qty
}

func manager(animal string) func(qty float64) float64 {
	switch animal {
	case "perro":
		return qty_food_perro

	case "gato":
		return qty_food_gato

	case "tarantula":
		return qty_food_tarantula

	case "hamster":
		return qty_food_hamster
	}
	return nil
}

const (
	perro     = "perro"
	gato      = "gato"
	tarantula = "tarantula"
	hamster   = "hamster"
)

func main() {
	op := manager(gato)
	res := op(3)
	fmt.Printf("\n La cantidad de alimento es: %.2f g \n", res)
}
