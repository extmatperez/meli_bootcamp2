package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	a1 := rand.Perm(100)
	a2 := rand.Perm(100)
	a3 := rand.Perm(100)

	//b1 := rand.Perm(1000)
	//b2 := rand.Perm(1000)
	//b3 := rand.Perm(1000)

	//c1 := rand.Perm(10000)
	//c2 := rand.Perm(10000)
	//c3 := rand.Perm(10000)

	channel1 := make(chan int64)
	go Burbuja(&a1, channel1)

	channel2 := make(chan int64)
	go Insercion(&a2, channel2)

	channel3 := make(chan int64)
	go Seleccion(&a3, channel3)

	fmt.Printf("Burbuja - Milisegundos: %v\n", <-channel1)
	fmt.Printf("Insercion - Milisegundos: %v\n", <-channel2)
	fmt.Printf("Seleccion - Milisegundos: %v\n", <-channel3)

}

func Insercion(variables *[]int, c chan int64) {
	inicio := time.Now()
	var auxiliar int
	for i := 1; i < len(*variables); i++ {
		auxiliar = (*variables)[i]
		for j := i - 1; j >= 0 && (*variables)[j] > auxiliar; j-- {
			(*variables)[j+1] = (*variables)[j]
			(*variables)[j] = auxiliar
		}
	}
	fin := time.Now()
	total := fin.Sub(inicio)
	c <- total.Microseconds()
}

func Burbuja(variables *[]int, c chan int64) {
	inicio := time.Now()
	var auxiliar int
	for i := 0; i < len(*variables); i++ {
		for j := 0; j < len(*variables); j++ {
			if (*variables)[i] > (*variables)[j] {
				auxiliar = (*variables)[i]
				(*variables)[i] = (*variables)[j]
				(*variables)[j] = auxiliar
			}
		}
	}
	fin := time.Now()
	total := fin.Sub(inicio)
	c <- total.Microseconds()
}

func Seleccion(variables *[]int, c chan int64) {
	inicio := time.Now()
	for i := 0; i < len(*variables)-1; i++ {
		min := i
		for j := i + 1; j < len(*variables); j++ {
			if (*variables)[j] < (*variables)[min] {
				min = j
			}
		}
		if i != min {
			aux := (*variables)[i]
			(*variables)[i] = (*variables)[min]
			(*variables)[min] = aux
		}
	}
	fin := time.Now()
	total := fin.Sub(inicio)
	c <- total.Microseconds()
}
