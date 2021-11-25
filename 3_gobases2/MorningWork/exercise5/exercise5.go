/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.


Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un error (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

*/
package main

import "fmt"

const (
	tarantula = "tarantula"
	hamster   = "hamster"
	perro     = "perro"
	gato      = "gato"
)

func animalPerro(numbersAnimal int) int {
	return numbersAnimal * 10
}
func animalGato(numbersAnimal int) int {
	return numbersAnimal * 5
}
func animalTarantula(numbersAnimal int) int {
	return numbersAnimal * 150
}
func animalHamster(numbersAnimal int) int {
	return numbersAnimal * 250
}

func errFunc(number int) int {
	return -9999
}

func animales(tipoAnimales string) (func(number int) int, string) {
	var message string = "OK"
	switch tipoAnimales {
	case tarantula:
		return animalTarantula, message
	case hamster:
		return animalHamster, message
	case perro:
		return animalPerro, message
	case gato:
		return animalGato, message
	default:
		message = "Funcion no definida"
		return errFunc, message
	}
}

func main() {
	animalElegido := "gato"
	countAnimal := 5
	animal, err := animales(animalElegido)
	var comidaPorAnimal = animal(countAnimal)

	if err != "OK" {
		fmt.Print("ERROR ANIMAL NO EXISTENTE \n")
	} else {
		fmt.Printf("comida por Animal: %d\n", comidaPorAnimal)
	}

}
