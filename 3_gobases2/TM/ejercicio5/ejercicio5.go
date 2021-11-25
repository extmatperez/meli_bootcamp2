package main

import (
	"errors"
	"fmt"
)

/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.


Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un error (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.


ejemplo:
const (
   perro = "perro"
   gato = "gato"
)

...

animalPerro, err := Animal(perro)
animalGato, err := Animal(gato)
...

var cantidad float64
cantidad += animalPerro(5)
cantidad += animalGato(8)
*/

func animal(animal string, quantity int) {
	function, err := funcTuUse(animal)

	if err != nil {
		panic("Wrong animal")
	}
	fmt.Printf("You need to buy %v kg \n", function(quantity))
}

func funcTuUse(animal string) (func(int) float64, error) {
	switch animal {
	case "perro":
		return Perro, nil
	case "gato":
		return Gato, nil
	case "hamster":
		return Hamster, nil
	case "tarantula":
		return Tarantula, nil
	default:
		return nil, errors.New("Wrong Animal")
	}
}

func Perro(quantity int) float64 {
	return float64(quantity) * 10
}
func Gato(quantity int) float64 {
	return float64(quantity) * 5
}
func Hamster(quantity int) float64 {
	return float64(quantity) * 0.25
}
func Tarantula(quantity int) float64 {
	return float64(quantity) * 0.15
}

func main() {

}
