package main

import (
	"fmt"
	"errors"
	"sort"
)

const (
	minimo = "minimo"
	promedio = "promedio"
	maximo = "maximo"
 )
 

func main() {
	pruebaOperacion(minimo)
	pruebaOperacion(promedio)
	pruebaOperacion(maximo)
	pruebaOperacion("otra distinta")

 }


func operacion (operacion string)  (func(notas... float64) float64,error) {
	
	switch operacion{
	case minimo:
		return getMinimo,nil
	case maximo:
		return getMaximo,nil
	case promedio:
		return getPromedio,nil
	default:
		return nil, errors.New("la operación no existe")
	}

}

func pruebaOperacion(operador string) {
	opFunc, err := operacion(operador)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El valor devuelto por la operacion %s es de %v\n", operador, 
		opFunc(2, 3, 3, 4, 1, 2, 4, 5))
	}
}




func getMinimo(notas... float64) (float64){
	cantidadNotas := len(notas)
	if(cantidadNotas == 0){
		return 0
		}else{
	 sort.Float64s(notas)

	return notas[0]
	}
}
		
func getMaximo(notas... float64) (float64){
	cantidadNotas := len(notas)
	if(cantidadNotas == 0){
		return 0
		}else{
	 sort.Float64s(notas)
	return notas[cantidadNotas-1]
	}
}
	
func getPromedio(notas... float64) (float64){
	cantidadNotas := len(notas)

	if(cantidadNotas == 0){
		return 0
		}else{
			sum := 0.0
			for _, nota := range notas {
				sum += nota
			}
			return sum/float64(cantidadNotas)
		}
}
	



