package main

import (
	"errors"
	"fmt"
)

const (
	perro = "perro"
	gato = "gato"
	hamster = "hamster"
	tarantula = "tarantula"
 )

var cantidades map[string]int;


func main() {
	// todos los pesos en gramos
	cantidades = map[string]int{
		perro: 10000,
		gato: 5000,
		hamster: 250,
		tarantula: 150,
	}
	 
	animalPerro, _ := Animal(perro)
	animalGato, _ := Animal(gato)
	animalInvalid, err := Animal("invalid")
	if (err != nil){
		fmt.Println("Error:", err)
	} else {
		animalInvalid(0)

		fmt.Println(animalPerro(5))
		fmt.Println(animalGato(2))
	}
}

func Animal(animal string) (func (cantidad int) int, error) {
	return calcularCantidad((animal))
}

func calcularCantidad(animal string) (func(cantidad int) int, error) {
	if (cantidades[animal] == 0) {
		return nil, errors.New("animal inválido")
	}
	return func(cantidad int) int {
		return cantidades[animal] * cantidad
	}, nil
}