package main

import (
	"errors"
	"fmt"
)

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
//    perro = "perro"
//    gato = "gato"
// )

// ...

// animalPerro, err := Animal(perro)
// animalGato, err := Animal(gato)

// ...

// var cantidad float64
// cantidad += animalPerro(5)
// cantidad += animalGato(8)

// Por perro necesitan 10 kg de alimento
// Por gato 5 kg
// Por cada Hamster 250 gramos.
// Por Tarántula 150 gramos.

func dogFood(dogCount int) int {
	foodQuantity := dogCount * 1000
	return foodQuantity
}

func catFood(catCount int) int {
	foodQuantity := catCount * 500
	return foodQuantity
}

func hamsterFood(hamsterCount int) int {
	foodQuantity := hamsterCount * 250
	return foodQuantity
}

func tarantulaFood(tarantulaCount int) int {
	foodQuantity := tarantulaCount * 150
	return foodQuantity
}

func calculateFood(animal string) (func(animalCount int) int, error) {

	switch animal {
	case "dog":
		return dogFood, errors.New("Error")
	case "cat":
		return catFood, errors.New("Error")
	case "hamster":
		return hamsterFood, errors.New("Error")
	case "tarantula":
		return tarantulaFood, errors.New("Error")
	}
	return nil, errors.New("the animal does not exist")

}

func main() {

	fmt.Println("Enter animal: ")
	var animal string
	fmt.Scanln(&animal)
	fmt.Println("Enter animal quantity:")
	var animal_count int
	fmt.Scanln(&animal_count)

	animalType, err := calculateFood(animal)
	foodNeeded := animalType(animal_count)

	if err == nil {
		fmt.Printf("For the %s it is needed: %d gr of food\n", animal, foodNeeded)

	} else {
		fmt.Printf("%s", err)

	}

}
