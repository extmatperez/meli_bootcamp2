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

func alimentoPerro(cant int) int {
	return 10000 * cant
}

func alimentoGato(cant int) int {
	return 5000 * cant
}

func alimentoHamster(cant int) int {
	return 250 * cant
}

func alimentoTarantula(cant int) int {
	return 150 * cant
}

func Animal(tipo string) (func(cant int) int, error) {

	switch tipo {
	case perro:
		return alimentoPerro, nil
	case gato:
		return alimentoGato, nil
	case hamster:
		return alimentoHamster, nil
	case tarantula:
		return alimentoTarantula, nil

	}
	return nil, errors.New("Error no existe funcion para el tipo de animal")
}
func main() {

	var cantidad int
	var tipoAnimal string

	fmt.Println("Ingresa el tipo de animal")
	fmt.Scanf("%s", &tipoAnimal)

	fmt.Printf("\nIngresa la cantidad de %v: ", tipoAnimal)
	fmt.Scanf("%d", &cantidad)

	funAnimal, err := Animal(tipoAnimal)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nNecesitas %d gramos de comida para %d %v \n", funAnimal(cantidad), cantidad, tipoAnimal)
	}

}
