package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarántula"
)

func main() {

	animalPerro, err := Animal(perro)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Se necesitarán %vkg de comida para %v perros\n", animalPerro(3), 3)
	}
	animalGato, err := Animal(gato)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Se necesitarán %vkg de comida para %v gatos\n", animalGato(3), 3)
	}
	animalHamster, err := Animal(hamster)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Se necesitarán %vkg de comida para %v hamsters\n", animalHamster(3), 3)
	}
	animalTarantula, err := Animal(tarantula)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Se necesitarán %vkg de comida para %v tarantulas\n", animalTarantula(3), 3)
	}
	animalPrueba, err := Animal("Prueba")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Se necesitarán %vkg de comida para %v tarantulas\n", animalPrueba(3), 3)
	}
}

func Animal(animal string) (func(cantAnimal int) float64, error) {
	switch animal {
	case perro:
		return calcularComidaPerro, nil
	case gato:
		return calcularComidaGato, nil
	case hamster:
		return calcularComidaHamster, nil
	case tarantula:
		return calcularComidaTarantula, nil
	default:
		return nil, errors.New("No existe el animal " + animal)
	}
}

func calcularComidaPerro(cant int) float64 {
	return 10.0 * float64(cant)
}
func calcularComidaGato(cant int) float64 {
	return 5.0 * float64(cant)
}
func calcularComidaHamster(cant int) float64 {
	return 0.25 * float64(cant)
}
func calcularComidaTarantula(cant int) float64 {
	return 0.150 * float64(cant)
}
