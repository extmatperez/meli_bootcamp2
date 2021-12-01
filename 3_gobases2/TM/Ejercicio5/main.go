/* Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen
tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

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

import (
	"fmt"
)

var cantPerros int

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	op := animal("tarantula")
	r := op(5)

	fmt.Println(r)

}

func animal(animal string) func(cantPerros int) string {
	switch animal {
	case "perro":
		return cantPerro

	case "gato":
		return cantGato
	case "hamster":
		return cantHamster
	case "tarantula":
		return cantTarantula
	default:
		fmt.Println("no existe la categoria")
	}
	return nil
}

func cantPerro(cantPerros int) string {

	var food int
	food = cantPerros * 10
	//texto := fmt.Sprintf("kg %d\n", food)

	return fmt.Sprintf("%d kg\n", food)

}

func cantGato(cantGatos int) string {

	var food int
	food = cantGatos * 5

	return fmt.Sprintf("%d kg\n", food)

}

func cantTarantula(cantTarantulas int) string {

	var food int
	food = cantTarantulas * 150

	return fmt.Sprintf("%d grs\n", food)

}

func cantHamster(cantHamsters int) string {

	var food int
	food = cantHamsters * 250

	return fmt.Sprintf("%d grs\n", food)

}
