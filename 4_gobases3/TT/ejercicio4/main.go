package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	// variable4 := rand.Perm(100000)
	ci := make(chan int)
	cb := make(chan int)
	cs := make(chan int)
	fmt.Println(variable1)

	go OrdenarPorInsercion(variable1, ci)
	go OrdenarPorBurbuja(variable1, cb)
	go OrdenarPorSeleccion(variable1, cs)

	fmt.Printf("\nTamaño 100:\nInsercion: %d\nBurbuja: %d\nSeleccion: %d\n", <-ci, <-cb, <-cs)

	go OrdenarPorInsercion(variable2, ci)
	go OrdenarPorBurbuja(variable2, cb)
	go OrdenarPorSeleccion(variable2, cs)
	fmt.Printf("\nTamaño 1000:\nInsercion: %d\nBurbuja: %d\nSeleccion: %d\n", <-ci, <-cb, <-cs)

	go OrdenarPorInsercion(variable3, ci)
	go OrdenarPorBurbuja(variable3, cb)
	go OrdenarPorSeleccion(variable3, cs)
	fmt.Printf("\nTamaño 10000:\nInsercion: %d\nBurbuja: %d\nSeleccion: %d\n", <-ci, <-cb, <-cs)

	// go OrdenarPorInsercion(variable4, ci)
	// go OrdenarPorBurbuja(variable4, cb)
	// go OrdenarPorSeleccion(variable4, cs)
	// fmt.Printf("\nTamaño 100000:\nInsercion: %d\nBurbuja: %d\nSeleccion: %d\n", <-ci, <-cb, <-cs)

}

func OrdenarPorInsercion(arr []int, c chan int) {
	inicio := time.Now()

	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	//fmt.Println("Insertion: ", arr)
	c <- int(time.Now().Sub(inicio).Microseconds())

}

func OrdenarPorBurbuja(arr []int, c chan int) {
	inicio := time.Now()

	for i := len(arr); i > 0; i-- {
		//The inner loop will first iterate through the full length
		//the next iteration will be through n-1
		// the next will be through n-2 and so on
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	//fmt.Println("Bubble: ", arr)

	c <- int(time.Now().Sub(inicio).Microseconds())
}

func OrdenarPorSeleccion(arr []int, c chan int) {
	inicio := time.Now()

	var n = len(arr)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	//fmt.Println("Selection: ", arr)

	c <- int(time.Now().Sub(inicio).Microseconds())
}
