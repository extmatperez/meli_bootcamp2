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

func perroFunc(cantidad_perros int) (alimento float64) {
	alimento = float64(cantidad_perros * 10)
	return
}

func gatoFunc(cantidad_gatos int) (alimento float64) {
	alimento = float64(cantidad_gatos * 5)
	return
}

func hamsterFunc(cantidad_hamsters int) (alimento float64) {
	alimento = float64(float64(cantidad_hamsters) * 0.25)
	return
}

func tarFunc(cantidad_tar int) (alimento float64) {
	alimento = float64(cantidad_tar) * 0.150
	return
}

func main() {

	cantidad, err := Animal(hamster)
	if err != nil {
		fmt.Println(err)
	} else {
		result := cantidad(5)
		fmt.Println("Se requieren ", result, " kg de alimento")
	}
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case perro:
		return perroFunc, nil
	case gato:
		return gatoFunc, nil
	case hamster:
		return hamsterFunc, nil
	case tarantula:
		return tarFunc, nil
	}

	return nil, errors.New("no existe el animal indicado")
}
