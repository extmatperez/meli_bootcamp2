package main

import (
	"errors"
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantua"
)

func perroFunc(cantidad int) int {
	return cantidad * 10000
}
func gatoFunc(cantidad int) int {
	return cantidad * 5000
}
func hamsterFunc(cantidad int) int {
	return cantidad * 250
}
func tarantulaFunc(cantidad int) int {
	return cantidad * 150
}

func Animal(animal string) (func(int) int, error) {
	switch animal {
	case perro:
		return perroFunc, nil
	case gato:
		return gatoFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	}
	return nil, errors.New("el animal" + animal + "no se encuentra registrado")
}

func main() {
	var cantidad float64
	animalPerro, err := Animal(perro)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		cantidad += float64(animalPerro(5))
	}
	animalGato, err := Animal(gato)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		cantidad += float64(animalGato(1))
	}
	animalDesconocido, err := Animal("asda")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		cantidad += float64(animalDesconocido(1))
	}
	fmt.Printf("La cantidad de alimento a comprar es: %.3f Kg\n", cantidad/1000)
}
