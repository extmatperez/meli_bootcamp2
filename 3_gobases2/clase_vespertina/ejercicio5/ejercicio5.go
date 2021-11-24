package main

import (
	"errors"
	"fmt"
)

func main() {
	var animal string
	fmt.Printf("Ingrese el animal a calcular:")
	fmt.Scanf("%s", &animal)
	fun, err := Animal(animal)
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		cant := fun(10)
		fmt.Printf("Se necesitan %v kgs.\n", cant)
	}
}

func Animal(animal string) (func(cantidad uint16) float64, error) {
	animales := map[string]float64{"perro": 10.0, "gato": 5.0, "hamster": (250.0 / 1000.0), "tarantula": (150.0 / 1000.0)}
	switch animal {
	case "perro", "gato", "hamster", "tarantula":
		fun := func(cantidad uint16) float64 {
			return float64(cantidad) * animales[animal]
		}
		return fun, nil
	default:
		return nil, errors.New("No existe el animal\n")
	}
}
