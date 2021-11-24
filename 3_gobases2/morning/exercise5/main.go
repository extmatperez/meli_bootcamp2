package main

import (
	"fmt"
)

const (
	perro     = "perro"
	gato      = "gato"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func sumDog(animalCantity int) float64 {
	return float64(animalCantity) * 10
}
func sumCat(animalCantity int) float64 {
	return float64(animalCantity) * 5
}
func sumHamster(animalCantity int) float64 {
	return float64(animalCantity) * 0.25
}
func sumTarantula(animalCantity int) float64 {
	return float64(animalCantity) * 0.15
}

func Animal(animalType string) (func(int) float64, error) {

	switch animalType {
	case perro:
		return sumDog, nil
	case gato:
		return sumCat, nil
	case hamster:
		return sumHamster, nil
	case tarantula:
		return sumTarantula, nil
	default:
		return nil, fmt.Errorf("El animal %s no existe", animalType)
	}
}

func main() {
	/*
		Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

		Por perro necesitan 10 kg de alimento
		Por gato 5 kg
		Por cada Hamster 250 gramos.
		Por Tarántula 150 gramos.

		Se solicita:
		Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un error (en caso que no exista el animal)
		Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
	*/

	flag := true
	var animalType string
	animalCantity := 0
	foodCantity := 0.0

	for flag == true {
		fmt.Printf("Ingrese el animal que desea calcular o 'salir' pata terminar: \n")
		fmt.Scanf("%s", &animalType)

		if animalType == "salir" {
			flag = false
			break
		} else {

			fmt.Printf("Ingrese la cantidad de animales: \n")
			fmt.Scanf("%d", &animalCantity)

			getTotal, err := Animal(animalType)

			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				foodCantity += getTotal(animalCantity)
			}
		}
	}
	fmt.Printf("El total de alimento necesario es: %.2f\n", foodCantity)
}
