package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ordenamientoInsercion(arreglo []int, canal chan string) {
	//	fmt.Printf("\nInsercion: %v \n", arreglo)
	inicio := time.Now()
	for i := 1; i < len(arreglo); i++ {
		clave := arreglo[i]
		j := i - 1
		//Comparar el valor selecionado con todos los valores anteriores
		for j >= 0 && arreglo[j] > clave {
			//Insertar el valor donde corresponda
			arreglo[j+1] = arreglo[j]
			j = j - 1
		}
		arreglo[j+1] = clave
	}
	fin := time.Now().Sub(inicio)
	fmt.Println("Insercion demoro: ", fin)
	//	fmt.Printf("\nFin Insercion: %v \n", arreglo)
	canal <- "fin"
}

func ordenamientoBurbuja(lista []int, canal chan string) {

	//	fmt.Printf("\nBurbuja: %v \n", lista)

	inicio := time.Now()
	n := 1
	l := len(lista)
	for n != 0 {
		n = 0
		//Recorrer la lista
		for i := 1; i < l; i++ {
			//Verificar si los dos valores estan ordenados
			if lista[i-1] > lista[i] {
				//Ordenar si es necesario
				temp := lista[i-1]
				lista[i-1] = lista[i]
				lista[i] = temp
				n = i
			}
		}
		l = n
	}
	fin := time.Now().Sub(inicio)
	fmt.Println("Burujar demoro: ", fin)
	//	fmt.Printf("\nFin Burbuja: %v \n", lista)
	canal <- "fin"
}

func ordenamientoSeleccion(a []int, canal chan string) {
	//	fmt.Printf("\nSeleccion: %v \n", a)
	inicio := time.Now()
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if i != min {
			aux := a[i]
			a[i] = a[min]
			a[min] = aux
		}
	}
	fin := time.Now().Sub(inicio)
	fmt.Println("Seleccion demoro: ", fin)

	canal <- "fin"
}

func main() {

	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	vb1 := make([]int, 100)
	vi1 := make([]int, 100)
	vs1 := make([]int, 100)

	copy(vb1, variable1)
	copy(vi1, variable1)
	copy(vs1, variable1)

	vb2 := make([]int, 1000)
	vi2 := make([]int, 1000)
	vs2 := make([]int, 1000)

	copy(vb2, variable2)
	copy(vi2, variable2)
	copy(vs2, variable2)

	vb3 := make([]int, 10000)
	vi3 := make([]int, 10000)
	vs3 := make([]int, 10000)

	copy(vb3, variable3)
	copy(vi3, variable3)
	copy(vs3, variable3)

	chanBurbuja := make(chan string)
	chanInsercion := make(chan string)
	chanSeleccion := make(chan string)

	chanBurbuja2 := make(chan string)
	chanInsercion2 := make(chan string)
	chanSeleccion2 := make(chan string)

	chanBurbuja3 := make(chan string)
	chanInsercion3 := make(chan string)
	chanSeleccion3 := make(chan string)

	fmt.Println("---------------------------------")

	go ordenamientoBurbuja(vb1, chanBurbuja)
	go ordenamientoInsercion(vi1, chanInsercion)
	go ordenamientoSeleccion(vs1, chanSeleccion)

	fmt.Printf("\nOrdenamiento1 burbuja: %s\n", <-chanBurbuja)
	fmt.Printf("\nOrdenamiento1 insercion: %s\n", <-chanInsercion)
	fmt.Printf("\nOrdenamiento1 seleccion: %s\n", <-chanSeleccion)

	fmt.Println("---------------------------------")

	go ordenamientoBurbuja(vb2, chanBurbuja2)
	go ordenamientoInsercion(vi2, chanInsercion2)
	go ordenamientoSeleccion(vs2, chanSeleccion2)

	fmt.Printf("\nOrdenamiento2 burbuja: %s\n", <-chanBurbuja2)
	fmt.Printf("\nOrdenamiento2 insercion: %s\n", <-chanInsercion2)
	fmt.Printf("\nOrdenamiento2 seleccion: %s\n", <-chanSeleccion2)

	fmt.Println("---------------------------------")

	go ordenamientoBurbuja(vb3, chanBurbuja3)
	go ordenamientoInsercion(vi3, chanInsercion3)
	go ordenamientoSeleccion(vs3, chanSeleccion3)

	fmt.Printf("\nOrdenamiento3 burbuja: %s\n", <-chanBurbuja3)
	fmt.Printf("\nOrdenamiento3 insercion: %s\n", <-chanInsercion3)
	fmt.Printf("\nOrdenamiento3 seleccion: %s\n", <-chanSeleccion3)

}
