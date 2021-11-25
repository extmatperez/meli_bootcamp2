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

	c := make(chan float64)
	
	go probarOrdenamiento(variable1, variable2, variable3, ordenamientoPorInsercion, c)
	fmt.Println("Ordenamiento por insercion: ", <-c)
	
	go probarOrdenamiento(variable1, variable2, variable3, ordenamientoPorBurbujas, c)
	fmt.Println("Ordenamiento por burbujas: ", <-c)

	go probarOrdenamiento(variable1, variable2, variable3, ordenamientoPorSeleccion, c)
	fmt.Println("Ordenamiento por seleccion: ", <-c)
}

func probarOrdenamiento(vector1 []int, vector2 []int, vector3 []int, orderer func([]int) []int, c chan float64) {
	inicio := time.Now()
	orderer(vector1)
	orderer(vector2)
	orderer(vector3)
	fin := time.Now()
	duracion := fin.Sub(inicio)

	c <- duracion.Seconds()
}

func ordenamientoPorInsercion(vector []int) []int {
	for i := 1; i < len(vector); i++ {
		j := i - 1
		for j >= 0 && vector[j] > vector[j+1] {
			vector[j], vector[j+1] = vector[j+1], vector[j]
			j--
		}
	}

	return vector
}

func ordenamientoPorBurbujas(vector []int) []int {
	for i := 0; i < len(vector); i++ {
		for j := 0; j < len(vector)-1; j++ {
			if vector[j] > vector[j+1] {
				vector[j], vector[j+1] = vector[j+1], vector[j]
			}
		}
	}

	return vector
}

func ordenamientoPorSeleccion(vector []int) []int {
	for i := 0; i < len(vector); i++ {
		min := i
		for j := i + 1; j < len(vector); j++ {
			if vector[j] < vector[min] {
				min = j
			}
		}
		vector[i], vector[min] = vector[min], vector[i]
	}
	return vector
}