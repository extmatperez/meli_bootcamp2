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
	var cantidad float64

	funcPerro, err := Animal(perro)
	funcGato, err := Animal(gato)
	funcHamster, err := Animal(hamster)
	funcTarantula, err := Animal(tarantula)

	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += funcPerro(1)
		cantidad += funcGato(1)
		cantidad += funcTarantula(1)
		cantidad += funcHamster(1)
	}
	fmt.Println("la cantidad de comida a comprar es:", cantidad)
}

func Animal(tipoAnimal string) (func(cantAnimales int) float64, error) {
	switch tipoAnimal {
	case perro:
		return ComidaPerro, nil
	case gato:
		return ComidaGato, nil
	case hamster:
		return ComidaHamster, nil
	case tarantula:
		return ComidaTarantula, nil
	default:
		return nil, errors.New("el animal no existe")
	}
}

func ComidaPerro(cantAnimales int) float64 {
	cantComida := float64(cantAnimales) * 10
	return cantComida
}

func ComidaGato(cantAnimales int) float64 {
	cantComida := float64(cantAnimales) * 5
	return cantComida
}

func ComidaHamster(cantAnimales int) float64 {
	cantComida := float64(cantAnimales) * 0.25
	return cantComida
}

func ComidaTarantula(cantAnimales int) float64 {
	cantComida := float64(cantAnimales) * 0.15
	return cantComida
}
