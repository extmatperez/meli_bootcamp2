package main

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
import (
	"fmt"
	"math/rand"
	"time"
)

func insertionsort(item []int) {
	items := item
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

func bubblesort(item []int) {
	items := item
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func selectionsort(item []int) {
	items := item
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

func proceso1(item1 []int, item2 []int, item3 []int, c chan int) {
	ini := time.Now()
	insertionsort(item1)
	insertionsort(item2)
	insertionsort(item3)
	fin := time.Now()
	tiempo := fin.Sub(ini)
	fmt.Printf("Tiempo total Insertion Sort: %.4v segundos\n", tiempo.Milliseconds())
	c <- 1
}

func proceso2(item1 []int, item2 []int, item3 []int, c chan int) {
	ini := time.Now()
	bubblesort(item1)
	bubblesort(item2)
	bubblesort(item3)
	fin := time.Now()
	tiempo := fin.Sub(ini)
	fmt.Printf("Tiempo total Bubble Sort: %.4v Milisegundos\n", tiempo.Milliseconds())
	c <- 2
}

func proceso3(item1 []int, item2 []int, item3 []int, c chan int) {
	ini := time.Now()
	selectionsort(item1)
	selectionsort(item2)
	selectionsort(item3)
	fin := time.Now()
	tiempo := fin.Sub(ini)
	fmt.Printf("Tiempo total Selection Sort: %.4v Milisegundos\n", tiempo.Milliseconds())
	c <- 3
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	go proceso1(variable1, variable2, variable3, a)
	go proceso2(variable1, variable2, variable3, b)
	go proceso3(variable1, variable2, variable3, c)
	<-a
	<-b
	<-c
}
