package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ordenamientoBurbujeo(vec []int, c chan []int) {
	var aux int
	ini := time.Now()
	for i := 0; i < len(vec); i++ {
		for j := 0; j < len(vec); j++ {
			if vec[i] > vec[j] {
				aux = vec[i]
				vec[i] = vec[j]
				vec[j] = aux
			}
		}
	}
	fin := time.Now()
	fmt.Println("El ordenamiento por burbujeo tardó", fin.Sub(ini).Seconds()*100, "milisegundos")
	c <- vec
}

func ordenamientoSeleccion(vec []int, c chan []int) {
	var aux int
	ini := time.Now()
	for i := 0; i < len(vec)-1; i++ {
		for j := i + 1; j < len(vec); j++ {
			if vec[i] > vec[j] {
				aux = vec[i]
				vec[i] = vec[j]
				vec[j] = aux
			}
		}
	}
	fin := time.Now()
	fmt.Println("El ordenamiento por selección tardó", fin.Sub(ini).Seconds()*100, "milisegundos")
	c <- vec
}

func ordenamientoInsercion(vec []int, c chan []int) {
	var aux int
	ini := time.Now()
	for i := 1; i < len(vec); i++ {
		aux = vec[i]
		for j := i - 1; j >= 0 && vec[j] > aux; j-- {
			vec[j+1] = vec[j]
			vec[j] = aux
		}
	}
	fin := time.Now()
	fmt.Println("El ordenamiento por inserción tardó", fin.Sub(ini).Seconds()*100, "milisegundos")
	c <- vec
}

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	c := make(chan []int)

	fmt.Println("Tiempos para ordenamiento de 100 elementos:")
	go ordenamientoBurbujeo(variable1, c)
	go ordenamientoInsercion(variable1, c)
	go ordenamientoSeleccion(variable1, c)
	for i := 0; i < 3; i++ {
		<-c
	}

	fmt.Println("Tiempos para ordenamiento de 1000 elementos:")
	go ordenamientoBurbujeo(variable2, c)
	go ordenamientoInsercion(variable2, c)
	go ordenamientoSeleccion(variable2, c)
	for i := 0; i < 3; i++ {
		<-c
	}

	fmt.Println("Tiempos para ordenamiento de 10000 elementos:")
	go ordenamientoBurbujeo(variable3, c)
	go ordenamientoInsercion(variable3, c)
	go ordenamientoSeleccion(variable3, c)
	for i := 0; i < 3; i++ {
		<-c
	}
}
