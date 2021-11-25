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

func main() {
	var cantidad float64 = 0.0
	CalcularAlimento, err := Animal(perro)
	if err == nil {
		cantidad += CalcularAlimento(4)
	} else {
		fmt.Printf("Error 1\n")
	}
	CalcularAlimento, err = Animal(gato)
	if err == nil {
		cantidad += CalcularAlimento(2)
	} else {
		fmt.Printf("Error 2\n")
	}
	CalcularAlimento, err = Animal(hamster)
	if err == nil {
		cantidad += CalcularAlimento(7)
	} else {
		fmt.Printf("Error 3\n")
	}
	CalcularAlimento, err = Animal(tarantula)
	if err == nil {
		cantidad += CalcularAlimento(10)
	} else {
		fmt.Printf("Error 4\n")
	}
	CalcularAlimento, err = Animal("invalido")
	if err == nil {
		cantidad += CalcularAlimento(10)
	} else {
		fmt.Printf("Error 5\n")
	}

	fmt.Printf("Cantidad total: %.2f kg\n", cantidad)

}

func Animal(tipo string) (func(int) float64, error) {
	switch tipo {
	case "perro":
		return AlimentoPerro, nil
	case "gato":
		return AlimentoGato, nil
	case "hamster":
		return AlimentoHamster, nil
	case "tarantula":
		return AlimentoTarantula, nil
	default:
		return nil, errors.New("Animal inválido")
	}
}

func AlimentoPerro(cantidad int) float64 {
	return float64(cantidad) * 10.0
}

func AlimentoGato(cantidad int) float64 {
	return float64(cantidad) * 5.0
}

func AlimentoHamster(cantidad int) float64 {
	return float64(cantidad) * 0.250
}

func AlimentoTarantula(cantidad int) float64 {
	return float64(cantidad) * 0.150
}
