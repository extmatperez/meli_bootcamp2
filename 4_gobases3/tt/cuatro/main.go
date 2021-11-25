package main

import (
	"fmt"
	"math/rand"
)

func swap(arr *[]int, i, j int) {
	aux := (*arr)[j]
	(*arr)[j] = (*arr)[i]
	(*arr)[i] = aux
}

func insertionSort(array *[]int) {
	for i := 1; i < len(*array); i++ { //itero sobre el array
		if (*array)[i] < (*array)[i-1] { //si el numero esta desordenado
			var j int
			aux := (*array)[i]
			for j = i - 1; j >= 0; j-- { //busco la posicion a insertar
				if (*array)[j] < aux { //corto el bucle una vez la encuentre
					break
				}
				(*array)[j+1] = (*array)[j]
			}
			//swap(&array, i, j) //inserto el elemento
			(*array)[j+1] = aux
		}
	}
}

func bubbleSort(array *[]int) {
	var swapped bool
	for i := 0; i < len(*array)-1; i++ {
		swapped = false
		for j := 0; j < len(*array)-i-1; j++ {
			if (*array)[j] > (*array)[j+1] {
				swap(array, j, j+1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func selectionSort(array *[]int) {
	for i := 0; i < len(*array)-1; i++ {
		min := i
		for j := i + 1; j < len(*array); j++ {
			if (*array)[j] < (*array)[min] {
				min = j
			}
			swap(array, min, i)
		}
	}
}

func main() {
	array1 := rand.Perm(10)
	array2 := rand.Perm(10)
	array3 := rand.Perm(10)

	fmt.Println(array1)
	insertionSort(&array1)
	fmt.Println(array1)

	fmt.Println(array2)
	bubbleSort(&array2)
	fmt.Println(array2)

	fmt.Println(array3)
	selectionSort(&array3)
	fmt.Println(array3)
}
