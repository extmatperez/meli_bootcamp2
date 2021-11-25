/*
Por perro necesitan 10 kg de alimento
Por gato 5 kg
Por cada Hamster 250 gramos.
Por Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una
función y un error (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
*/

package main

import "fmt"

var cantPerros int

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func animal(animal string) {
	switch animal {
	case perro:
		fmt.Println("defina una cantidad de perros")
		fmt.Scanf("%d", &cantPerros)
		cantPerro(cantPerros)
		/*	case gato:
				cantGato()
			case hamster:
				cantHamster()
			case tarantula:
				cantTarantula()*/
	}

	return cantAlimento()

}

func cantPerro(cantPerros int) int {

	var food int
	food = cantPerros * 10

	return food

}

func cantGato(cantGatos int) int {

	var food int
	food = cantGatos * 5

	return food

}

func cantTarantula(cantTarantulas int) int {

	var food int
	food = cantTarantulas * 150

	return food

}

func cantHamster(cantHamsters int) int {

	var food int
	food = cantHamsters * 250

	return food

}

func main() {
	animalPerro, err := animal(perro)
	//animalGato, err := animal(gato)

	/*var cantidad float64
	cantidad += animalPerro(5)
	cantidad += animalGato(8)*/

}
