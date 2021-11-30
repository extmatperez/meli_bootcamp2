package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(arr []int, c chan float64) {
	ini := time.Now()
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
				i = j
			}
		}
	}
	fin := time.Now()
	c <- fin.Sub(ini).Seconds()
}
func bubbleSort(arr []int, c chan float64) {
	ini := time.Now()
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fin := time.Now()
	c <- fin.Sub(ini).Seconds()
}
func selectionSort(arr []int, c chan float64) {

	ini := time.Now()
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fin := time.Now()
	c <- fin.Sub(ini).Seconds()
}

func main() {

	variables := [][]int{rand.Perm(100), rand.Perm(1000), rand.Perm(10000) /* , rand.Perm(100000) */}

	chInsertion := make(chan float64)
	chBubble := make(chan float64)
	chSelection := make(chan float64)

	fmt.Printf("\n%-15s%-15s%-15s%-15s\n", "Cantidad", "Insertion", "Bubble", "Selection")
	fmt.Println("-------------------------------------------------------")

	for i := 0; i < len(variables); i++ {
		var arr = variables[i]
		var copia = make([]int, len(arr))

		copy(copia, arr)
		go insertionSort(copia, chInsertion)
		copy(copia, arr)
		go bubbleSort(copia, chBubble)
		copy(copia, arr)
		go selectionSort(copia, chSelection)

		timeInserion := <-chInsertion
		timeBubble := <-chBubble
		timeSelection := <-chSelection

		fmt.Printf("%-15d%-15.8f%-15.8f%-15.8f\n", len(variables[i]), timeInserion, timeBubble, timeSelection)
		fmt.Println("-------------------------------------------------------")
		fmt.Println()
	}
}

/*
	Una empresa de sistemas requiere analizar qué algoritmos de ordenamiento utilizar para sus servicios.
	Para ellos se requiere instanciar 3 arreglos con valores aleatorios desordenados
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
	Por último debo medir el tiempo de cada uno y mostrar en pantalla el resultado, para saber qué ordenamiento fue mejor para cada arreglo
*/
