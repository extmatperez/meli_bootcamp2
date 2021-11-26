/*
Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados:
un arreglo de números enteros con 100 valores
un arreglo de números enteros con 1000 valores
un arreglo de números enteros con 10000 valores

Para instanciar las variables utilizar rand
package main

import (
   "math/rand"
)

func main() {
   variable1 := rand.Perm(100)
   variable2 := rand.Perm(1000)
   variable3 := rand.Perm(10000)
}

Se debe realizar el ordenamiento de cada una por:
Ordenamiento por inserción
Ordenamiento por burbuja
Ordenamiento por selección

Una go routine por cada ejecución de ordenamiento
Debo esperar a que terminen los ordenamientos de 100 números para seguir el de 1000 y después el de 10000.
Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento
fue mejor para cada arreglo
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ordenamientoInsercion(c chan float64, variables []int) {

	inicio := time.Now()
	var aux int

	for i := 1; i < len(variables); i++ {

		aux = (variables)[i]

		for j := i - 1; j >= 0 && (variables)[j] > aux; j-- {
			(variables)[j+1] = (variables)[j]
			(variables)[j] = aux
		}
	}

	fin := time.Now()
	total := fin.Sub(inicio)

	c <- total.Seconds()
}

func ordenamientoBurbuja(c chan float64, variables []int) {

	inicio := time.Now()
	var aux int

	for i := 0; i < len(variables); i++ {

		for j := 0; j < len(variables); j++ {

			if (variables)[i] > (variables)[j] {
				aux = (variables)[i]
				(variables)[i] = (variables)[j]
				(variables)[j] = aux
			}
		}
	}

	fin := time.Now()
	total := fin.Sub(inicio)

	c <- total.Seconds()
}

func ordenamientoSeleccion(c chan float64, variables []int) {

	inicio := time.Now()

	for i := 0; i < len(variables)-1; i++ {

		min := i

		for j := i + 1; j < len(variables); j++ {

			if (variables)[j] < (variables)[min] {
				min = j
			}
		}

		if i != min {
			aux := (variables)[i]
			(variables)[i] = (variables)[min]
			(variables)[min] = aux
		}
	}

	fin := time.Now()
	total := fin.Sub(inicio)

	c <- total.Seconds()
}

func main() {

	cVariableInsercion1 := make(chan float64)
	go ordenamientoInsercion(cVariableInsercion1, rand.Perm(100))

	cVariableInsercion2 := make(chan float64)
	go ordenamientoInsercion(cVariableInsercion2, rand.Perm(100))

	cVariableInsercion3 := make(chan float64)
	go ordenamientoInsercion(cVariableInsercion3, rand.Perm(100))

	cVariableBurbuja1 := make(chan float64)
	go ordenamientoBurbuja(cVariableBurbuja1, rand.Perm(1000))

	cVariableBurbuja2 := make(chan float64)
	go ordenamientoBurbuja(cVariableBurbuja2, rand.Perm(1000))

	cVariableBurbuja3 := make(chan float64)
	go ordenamientoBurbuja(cVariableBurbuja3, rand.Perm(1000))

	cVariableSeleccion1 := make(chan float64)
	go ordenamientoSeleccion(cVariableSeleccion1, rand.Perm(10000))

	cVariableSeleccion2 := make(chan float64)
	go ordenamientoSeleccion(cVariableSeleccion2, rand.Perm(10000))

	cVariableSeleccion3 := make(chan float64)
	go ordenamientoSeleccion(cVariableSeleccion3, rand.Perm(10000))

	fmt.Printf("Inserción 100 enteros: %f\n", <-cVariableInsercion1)
	fmt.Printf("Inserción 1000 enteros: %f\n", <-cVariableInsercion2)
	fmt.Printf("Inserción 10000 enteros: %f\n", <-cVariableInsercion3)

	fmt.Println("")

	fmt.Printf("Burbuja 100 enteros: %f\n", <-cVariableBurbuja1)
	fmt.Printf("Burbuja 1000 enteros: %f\n", <-cVariableBurbuja2)
	fmt.Printf("Burbuja 10000 enteros: %f\n", <-cVariableBurbuja3)

	fmt.Println("")

	fmt.Printf("Selección 100 enteros: %f\n", <-cVariableSeleccion1)
	fmt.Printf("Selección 1000 enteros: %f\n", <-cVariableSeleccion2)
	fmt.Printf("Selección 10000 enteros: %f\n", <-cVariableSeleccion3)
}
