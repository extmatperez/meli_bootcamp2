package main

import "fmt"

var i string

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func calculation() {

	fmt.Println("¿Qué tipo de cálculo desea realizar?\n1.Mínimo\n2.Máximo\n3.Promedio")
	fmt.Scanf("%s", &i)
	operation()
}

func operation(operacion string){
	var notas []int
	notas = append(notas, 5, 2, 10)

	switch (operacion) {
	case minimo: for k:= 0 ; k<len(notas); k++ {
		if (k < notas[k+1] {
			j := notas[k+1]
		
		}
	}
	case promedio:
	case maximo:
	default:
		fmt.Println("El cálculo no está definido")
	}

	}


func main() {
	
	 minFunc, err := operacion(minimo)
	 promFunc, err := operacion(promedio)
	 maxFunc, err := operacion(maximo)

	 
	 valorMinimo := minFunc(2,3,3,4,1,2,4,5)
	 valorPromedio := promFunc(2,3,3,4,1,2,4,5)
	 valorMaximo := maxFunc(2,3,3,4,1,2,4,5)

	 calculation()

}
