package main

import (
	"errors"
	"fmt"
)

func main(){

	const (
		perro = "perro"
		gato = "gato"
		spider = "tarantula"
	)
	var quantity float64
	animalDog, err := animal_func(spider)

	if err != nil {
		fmt.Println(err)
	} else {
		quantity += animalDog(6)
		fmt.Println("La cantida de comida es: ", quantity, "Kg")
	}

}

func animal_func(animal string) (func(quantity float64) float64, error) {
	var (
		exist = []string{ "tarantula", "hamster", "perro", "gato"}
		check = false
	)

	for _, i := range exist {
		if animal == i {
			check = true
		}
	}
	if check == true {
		switch animal {
		case "tarantula":
			return tarantula, nil
		case "perro":
			return dog, nil
		case "gato":
			return cat, nil
		case "hamster":
			return hamster, nil
		}
	}
	return nil, errors.New("No existe ese animal")
}
func tarantula(quantity float64) float64 {
	totalFood := quantity * 0.15

	return totalFood
}
func dog(quantity float64) float64 {
	totalFood := quantity * 10

	return totalFood
}
func cat(quantity float64) float64 {
	totalFood := quantity * 5

	return totalFood
}
func hamster(quantity float64) float64 {
	totalFood := quantity * 0.25

	return totalFood
}