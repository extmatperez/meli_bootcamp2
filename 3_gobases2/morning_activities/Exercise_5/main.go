/* Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos,
pero se espera que puedan haber muchos más animales que refugiar.

Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.


Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un error
(en caso que no exista el animal) Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado. */

package main

import (
	"errors"
	"fmt"
)

func dog_food(quantity int) float64 {
	return float64(quantity) * 10
}

func cat_food(quantity int) float64 {
	return float64(quantity) * 5
}

func hamster_food(quantity int) float64 {
	return float64(quantity) * 0.250
}

func tarantula_food(quantity int) float64 {
	return float64(quantity) * 0.150
}

func food(animal string) (func(int) float64, error) {
	switch animal {
	case "dog":
		return dog_food, nil
	case "cat":
		return cat_food, nil
	case "hamster":
		return hamster_food, nil
	case "tarantula":
		return tarantula_food, nil
	default:
		return nil, errors.New("This animal doesn't exist in the list, try with another!")
	}
}

func ex5(animal string, quantity int) string {
	function, err := food(animal)

	if err != nil {
		panic("This animal doesn't exist in the list, please try with another one!")
	}

	fmt.Printf("Buy %.2f Kg of food to your "+animal+"! \n", function(quantity))
	return fmt.Sprintf("Buy %.2f Kg of food. \n", function(quantity))
}

func main() {
	ex5("dog", 4)
	ex5("cat", 4)
	ex5("hamster", 3)
	ex5("tarantula", 7)

}
