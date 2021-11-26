package main

import (
	"fmt"
	"errors"
	
)

const (
	perro = "perro"
	gato = "gato"
	tarantula = "tarantula"
	hamster = "hamster"
 )
 
 

func main() {

	var cantidad float64
	cantidad += orquestadorAnimal(perro,1)
	cantidad += orquestadorAnimal(gato,2)
	cantidad += orquestadorAnimal(hamster,4)

	fmt.Printf("El total de alimento es de %f kg\n",  cantidad)
 }


func orquestadorAnimal(animal string, cantidad float64) float64{

	animalOperador, err := Animal(animal)

	if err != nil {
		return -1
	} else {
		return animalOperador(cantidad)
	}
}

func Animal (tipoAnimal string)  (func(cantidad float64) float64,error) {
	
	switch tipoAnimal{
	case perro:
		return AnimalPerro,nil
	case gato:
		return AnimalGato,nil
	case hamster:
		return AnimalHamster,nil
	case tarantula:
		return AnimalTarantula,nil
	default:
		return nil, errors.New("no existe el animal")
	}

}






func AnimalPerro(cantidad float64) float64 {
	return cantidad * 10
}
		
func AnimalGato(cantidad float64) float64 {
	return cantidad * 5
}
func AnimalHamster(cantidad float64) float64 {
	return cantidad * 0.250
}
	
func AnimalTarantula(cantidad float64) float64 {
	return cantidad * 0.150
}


