package main

import (
	"fmt"
)

const (
	perro = "perro"
	gato  = "gato"
)

func main() {
	var cantidad float64

	err, AnimalPerro := Animal(perro)

	if err != nil {
		fmt.Errorf("There was a error %w", err)
	} else {

		cantidad += AnimalPerro(5)

		fmt.Println(cantidad)
	}

	err2, AnimalGato := Animal(gato)

	if err2 != nil {
		fmt.Errorf("There was a error %w", err2)
	} else {
		var cantidad float64
		cantidad += AnimalGato(8)

		fmt.Println(cantidad)
	}

	fmt.Println(cantidad)
}

func GetFoodByKg(value string) (float64, error) {

	var result float64
	switch value {
	case "dog":
		result = 10
	case "cat":
		result = 5
	case "hamster":
		result = 0.250
	case "tarantula":
		result = .150
	default:
		return 0.00, fmt.Errorf("There was a error")
	}
	return result, nil
}
func Animal(animalName string) (error, func(value float64) float64) {

	switch animalName {
	case "dog":
		return nil, AnimalPerro
	case "cat":
		return nil, AnimalGato
	case "hamster":
		return nil, AnimalHamster
	case "tarantula":
		return nil, AnimalTarantula
	default:
		return fmt.Errorf("There was a error"), nil
	}

}

func AnimalPerro(cantFood float64) float64 {
	return 10 * cantFood
}

func AnimalGato(cantFood float64) float64 {
	return 5 * cantFood
}

func AnimalHamster(cantFood float64) float64 {
	return 0.250 * cantFood
}

func AnimalTarantula(cantFood float64) float64 {
	return 0.150 * cantFood
}

type Dog struct {
	weight float64
	name   string
}
type Cat struct {
	weight float64
	name   string
}
