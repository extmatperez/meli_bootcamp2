package main

import (
	"errors"
	"fmt"
)

// Animals
const (
	DOG       = "perro"
	CAT       = "gato"
	HAMSTER   = "hamster"
	TARANTULA = "tarantula"
)

// Animals food (gr)
const (
	DOG_FOOD       = 10000
	CAT_FOOD       = 5000
	HAMSTER_FOOD   = 250
	TARANTULA_FOOD = 150
)

func main() {
	var amount float64

	dogAnimal, errDog := getAnimalFunc(DOG)
	catAnimal, errCat := getAnimalFunc(CAT)
	hamsterAnimal, errHamster := getAnimalFunc(HAMSTER)
	tarantulaAnimal, errTarantula := getAnimalFunc(TARANTULA)

	if errDog != nil {
		fmt.Println(errDog)
	} else {
		amount += dogAnimal(1)
	}

	if errCat != nil {
		fmt.Println(errCat)
	} else {
		amount += catAnimal(1)
	}

	if errHamster != nil {
		fmt.Println(errHamster)
	} else {
		amount += hamsterAnimal(1)
	}

	if errTarantula != nil {
		fmt.Println(errTarantula)
	} else {
		amount += tarantulaAnimal(1)
	}

	fmt.Printf("La cantidad de comida que se necesita es: %.2f gr\n", amount)
}

func getAnimalFunc(animal string) (func(amount int) float64, error) {
	switch animal {
	case DOG:
		return dogFood, nil
	case CAT:
		return catFood, nil
	case HAMSTER:
		return hamsterFood, nil
	case TARANTULA:
		return tarantulaFood, nil
	default:
		errorMsg := "el animal '" + animal + "' no esta definido"
		return func(amount int) float64 { return 0.0 }, errors.New(errorMsg)
	}
}

func dogFood(amount int) float64 {
	return float64(DOG_FOOD * amount)
}

func catFood(amount int) float64 {
	return float64(CAT_FOOD * amount)
}

func hamsterFood(amount int) float64 {
	return float64(HAMSTER_FOOD * amount)
}

func tarantulaFood(amount int) float64 {
	return float64(TARANTULA_FOOD * amount)
}
