// ¿Cuantos kg de alimento comprar?
// Por perro 10 kg, por gato 5 kg, por hamster 250 gr (0.25 kg) y tarantula 150 gr (0.15 kg)
// Funcion Animal tipo func Animal(nombre string) calculo(nombre string, peso float64) float64, error
// Funcion calculo(nombre string, peso float64) float64

package main

import (
	"errors"
	"fmt"
)

const (
	Perro     = "perro"
	Gato      = "gato"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

func Animal(nombre string) (func(qty int) float64, error) {
	switch nombre {
	case Perro:
		return calculo_alimento_perro, nil
	case Gato:
		return calculo_alimento_gato, nil
	case Hamster:
		return calculo_alimento_hamster, nil
	case Tarantula:
		return calculo_alimento_tarantula, nil
	default:
		return func(qty int) float64 { return 0.0 }, errors.New("No existe ese animal.")
	}
}

func calculo_alimento_perro(qty int) float64 {
	return (float64)(10000 * qty)
}

func calculo_alimento_gato(qty int) float64 {
	return (float64)(5000 * qty)
}

func calculo_alimento_hamster(qty int) float64 {
	return (float64)(250 * qty)
}

func calculo_alimento_tarantula(qty int) float64 {
	return (float64)(150 * qty)
}

func main() {
	animalPerro, err1 := Animal(Perro)
	animalGato, err2 := Animal(Gato)
	animalHamster, err3 := Animal(Hamster)
	animalTarantula, err4 := Animal(Tarantula)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("\nLa cantidad de comida para perros es %5.2f gramos", animalPerro(8))
	}

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("\nLa cantidad de comida para gatos es %5.2f gramos", animalGato(4))
	}

	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Printf("\nLa cantidad de comida para hamster es %5.2f gramos", animalHamster(14))
	}

	if err4 != nil {
		fmt.Println((err4))
	} else {
		fmt.Printf("\nLa cantidad de comida para tarantula es %5.2f gramos", animalTarantula(2))
	}
}
