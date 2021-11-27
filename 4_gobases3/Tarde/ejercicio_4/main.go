package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertion(vectorOri []int, c chan int) {
	vector := vectorOri
	for i := 1; i < len(vector); i++ {
		//clave = *(vector+i)
		j := i

		for j > 0 {
			if vector[j-1] > vector[j] {
				vector[j-1], vector[j] = vector[j], vector[j-1]
			}
			j = j - 1
		}

	}
	c <- 1
}

func bubble(vectorOri []int, c chan int) {
	vector := vectorOri
	n := len(vector)
	sorted := false

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if vector[i] > vector[i+1] {
				vector[i+1], vector[i] = vector[i], vector[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	c <- 2

}

func selection(vectorOri []int, c chan int) {
	vector := vectorOri
	var n = len(vector)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if vector[j] < vector[minIdx] {
				minIdx = j
			}
		}
		vector[i], vector[minIdx] = vector[minIdx], vector[i]
	}
	c <- 3
}

func main() {

	variable1 := rand.Perm(9100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	ini := time.Now()
	go insertion(variable1, c1)
	go insertion(variable2, c1)
	go insertion(variable3, c1)
	<-c1
	<-c1
	<-c1
	fin := time.Now()
	tiempo := fin.Sub(ini)
	fmt.Printf("Tiempo total Insercion es: %v ms\n", tiempo.Milliseconds())

	fmt.Println("\n---------")

	ini2 := time.Now()
	go bubble(variable1, c2)
	go bubble(variable2, c2)
	go bubble(variable3, c2)
	<-c2
	<-c2
	<-c2
	fin2 := time.Now()
	tiempo2 := fin2.Sub(ini2)
	fmt.Printf("Tiempo total Burbuja es: %v \n", tiempo2)

	fmt.Println("\n---------")

	ini3 := time.Now()
	go selection(variable1, c3)
	go selection(variable2, c3)
	go selection(variable3, c3)
	<-c3
	<-c3
	<-c3
	fin3 := time.Now()
	tiempo3 := fin3.Sub(ini3)
	fmt.Printf("Tiempo total seleccion es: %v ms\n", tiempo3.Milliseconds())

}
