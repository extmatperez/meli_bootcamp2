package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// variable4 := rand.Perm(100000)
	ci := make(chan float64)
	cb := make(chan float64)
	cs := make(chan float64)

	go OrdenarPorInsercion(100, ci)
	go OrdenarPorBurbuja(100, cb)
	go OrdenarPorSeleccion(100, cs)

	fmt.Printf("\nTamaño 100:\nInsercion: %f\nBurbuja: %f\nSeleccion: %f\n", <-ci, <-cb, <-cs)

	go OrdenarPorInsercion(1000, ci)
	go OrdenarPorBurbuja(1000, cb)
	go OrdenarPorSeleccion(1000, cs)
	fmt.Printf("\nTamaño 1000:\nInsercion: %f\nBurbuja: %f\nSeleccion: %f\n", <-ci, <-cb, <-cs)

	go OrdenarPorInsercion(10000, ci)
	go OrdenarPorBurbuja(10000, cb)
	go OrdenarPorSeleccion(10000, cs)
	fmt.Printf("\nTamaño 10000:\nInsercion: %f\nBurbuja: %f\nSeleccion: %f\n", <-ci, <-cb, <-cs)

	// go OrdenarPorInsercion(variable4, ci)
	// go OrdenarPorBurbuja(variable4, cb)
	// go OrdenarPorSeleccion(variable4, cs)
	// fmt.Printf("\nTamaño 100000:\nInsercion: %d\nBurbuja: %d\nSeleccion: %d\n", <-ci, <-cb, <-cs)

}

func OrdenarPorInsercion(size int, c chan float64) {
	inicio := time.Now()
	for iter := 0; iter < 100; iter++ {
		arr := rand.Perm(size)

		for i := 0; i < len(arr); i++ {
			for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}

	//fmt.Println("Insertion: ", arr)
	c <- time.Now().Sub(inicio).Seconds() / 100.0

}

func OrdenarPorBurbuja(size int, c chan float64) {
	inicio := time.Now()

	for iter := 0; iter < 100; iter++ {
		arr := rand.Perm(size)
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
	}
	//fmt.Println("Bubble: ", arr)

	c <- time.Now().Sub(inicio).Seconds() / 100.0
}

func OrdenarPorSeleccion(size int, c chan float64) {
	inicio := time.Now()
	for iter := 0; iter < 100; iter++ {
		arr := rand.Perm(size)

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
	}
	//fmt.Println("Selection: ", arr)

	c <- time.Now().Sub(inicio).Seconds() / 100.0
}
