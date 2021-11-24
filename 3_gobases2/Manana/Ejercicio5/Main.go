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

func cuantoPerro(cantidad int) float64 {
	return float64(cantidad * 10)
}

func cuantoGato(cantidad int) float64 {
	return float64(cantidad * 5)
}

//Hamster 250 tarantula 150
func cuantoHamster(cantidad int) float64 {
	return float64(cantidad) * 0.25
}

func cuantoTarantula(cantidad int) float64 {
	return float64(cantidad) * 0.15
}

func Animal(animal string) (func(int) float64, error) {

	switch animal {
	case perro:
		return cuantoPerro, nil
	case gato:
		return cuantoGato, nil
	case hamster:
		return cuantoHamster, nil
	case tarantula:
		return cuantoTarantula, nil
	default:
		return nil, errors.New("animal no valido")
	}
}

func main() {

	fmt.Println("Bienvenidos al ejercicio 5")

	var consumoTotal float64
	consumoTotal = 0.0

	fmt.Println("Hasta hora tenemos un consumo total de", consumoTotal)
	fmt.Println("CASO 1 con perro")
	funcPerro, err := Animal("perro")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Con 3 perros se consume", funcPerro(3))
		consumoTotal += funcPerro(3)
		fmt.Println("El consumo total va", consumoTotal)
	}

	fmt.Println("CASO 2 con tarantula")
	funcTarantula, err := Animal("tarantula")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Con 2 tarantulas se consume", funcTarantula(2))
		consumoTotal += funcTarantula(2)
		fmt.Println("El consumo total va", consumoTotal)
	}

	fmt.Println("CASO 3 con cocodrilo")
	funcCocodrilo, err := Animal("cocodrilo")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Con 10 cocodrilos se consume", funcCocodrilo(10))
		consumoTotal += funcCocodrilo(10)
		fmt.Println("El consumo total va", consumoTotal)
	}

}
