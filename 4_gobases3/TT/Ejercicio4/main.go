package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Algoritmo struct {
	Nombre string
	Tiempo time.Duration
}

func burbuja(lista []int, c chan Algoritmo) {
	ini := time.Now()
	aux := 0
	for i := 0; i < len(lista)-1; i++ {
		for j := 0; j < len(lista)-i-1; j++ {
			if lista[j+1] < lista[j] {
				aux = lista[j+1]
				lista[j+1] = lista[j]
				lista[j] = aux
			}
		}
	}
	fin := time.Now()
	tiempo := fin.Sub(ini)
	valor := Algoritmo{"burbuja", tiempo}
	c <- valor
}

func insercion(lista []int, c chan Algoritmo) {
	ini := time.Now()
	aux := 0
	j := 0
	for i := 1; i < len(lista); i++ {
		aux = lista[i]
		j = i - 1
		for j >= 0 && aux < lista[j] {
			lista[j+1] = lista[j]
			j--
		}
		lista[j+1] = aux
	}
	fin := time.Now()
	tiempo := fin.Sub(ini)
	valor := Algoritmo{"insercion", tiempo}
	c <- valor
}

func seleccion(lista []int, c chan Algoritmo) {
	ini := time.Now()
	var menor, pos, tmp int
	for i := 0; i < len(lista)-1; i++ {
		menor = lista[i]
		pos = i
		for j := i + 1; j < len(lista); j++ {
			if lista[j] < menor {
				menor = lista[j]
				pos = j
			}
		}
		if pos != i {
			tmp = lista[i]
			lista[i] = lista[pos]
			lista[pos] = tmp
		}
	}
	fin := time.Now()
	tiempo := fin.Sub(ini)
	valor := Algoritmo{"seleccion", tiempo}
	c <- valor
}

func main() {

	c100 := make(chan Algoritmo)
	c1000 := make(chan Algoritmo)
	c10000 := make(chan Algoritmo)

	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	var copia1 []int
	var copia2 []int
	var copia3 []int
	//array100 := []float64{}
	//array1000 := []float64{}
	//array10000 := []float64{}
	copy(copia1, variable1)
	copy(copia2, variable2)
	copy(copia3, variable3)

	go insercion(copia1, c100)
	go burbuja(copia1, c100)
	go seleccion(copia1, c100)

	for i := 0; i < 3; i++ {
		variable := <-c100
		fmt.Println(variable)
	}
	fmt.Println(variable1)

	go insercion(copia2, c1000)
	go burbuja(copia2, c1000)
	go seleccion(copia2, c1000)

	for i := 0; i < 3; i++ {
		variable := <-c1000
		fmt.Println(variable)
	}

	go insercion(copia3, c10000)
	go burbuja(copia3, c10000)
	go seleccion(copia3, c10000)

	for i := 0; i < 3; i++ {
		variable := <-c10000
		fmt.Println(variable)
	}

}
