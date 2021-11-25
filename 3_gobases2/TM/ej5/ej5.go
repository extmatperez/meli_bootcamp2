package ej5

import (
	"errors"
	"fmt"
)

const (
	tarantula = "Tar"
	hamster   = "Ham"
	perro     = "Perro"
	gato      = "Gato"
)

func Ej5(animal string, quantity int) string {
	fmt.Println("")
	function, err := funcToUse(animal)

	if err != nil {
		panic("You put the wrong animal!")
	}

	fmt.Printf("You'll need to buy %.2f kg\n", function(quantity))
	return fmt.Sprintf("You'll need to buy %f kg", function(quantity))
}

func funcToUse(animal string) (func(int) float64, error) {
	switch animal {
	case tarantula:
		return Tar, nil
	case hamster:
		return Ham, nil
	case perro:
		return Perro, nil
	case gato:
		return Gato, nil
	default:
		return nil, errors.New("no tenemos este animal en el refugio")
	}
}

func Tar(quantity int) float64 {
	return float64(quantity) * 0.15
}

func Ham(quantity int) float64 {
	return float64(quantity) * 0.25
}

func Perro(quantity int) float64 {
	return float64(quantity) * 10
}

func Gato(quantity int) float64 {
	return float64(quantity) * 5
}
