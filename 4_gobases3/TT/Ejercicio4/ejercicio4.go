package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr1 := rand.Perm(100)
	arr2 := rand.Perm(1000)
	arr3 := rand.Perm(10000)

	insertionSortChanArr := make(chan int64)
	selectionSortChanArr := make(chan int64)
	bubbleSortChanArr := make(chan int64)

	copy1 := make([]int, len(arr1))
	copy2 := make([]int, len(arr1))
	copy3 := make([]int, len(arr1))
	copy(copy1, arr1)
	copy(copy2, arr1)
	copy(copy3, arr1)

	go InsertionSort(copy1, insertionSortChanArr)
	go SelectionSort(copy2, selectionSortChanArr)
	go BubbleSort(copy3, bubbleSortChanArr)

	fmt.Println("=================")
	fmt.Println("Resultados 'arr1'")
	fmt.Println("=================")
	fmt.Printf("Ordenamiento por inserción: %d microsegundos\n", <-insertionSortChanArr)
	fmt.Printf("Ordenamiento por selección: %d microsegundos\n", <-selectionSortChanArr)
	fmt.Printf("Ordenamiento por burbuja: %d microsegundos\n", <-bubbleSortChanArr)

	copy1 = make([]int, len(arr2))
	copy2 = make([]int, len(arr2))
	copy3 = make([]int, len(arr2))
	copy(copy1, arr2)
	copy(copy2, arr2)
	copy(copy3, arr2)

	go InsertionSort(copy1, insertionSortChanArr)
	go SelectionSort(copy2, selectionSortChanArr)
	go BubbleSort(copy3, bubbleSortChanArr)

	fmt.Println()
	fmt.Println("=================")
	fmt.Println("Resultados 'arr2'")
	fmt.Println("=================")
	fmt.Printf("Ordenamiento por inserción: %d microsegundos\n", <-insertionSortChanArr)
	fmt.Printf("Ordenamiento por selección: %d microsegundos\n", <-selectionSortChanArr)
	fmt.Printf("Ordenamiento por burbuja: %d microsegundos\n", <-bubbleSortChanArr)

	copy1 = make([]int, len(arr3))
	copy2 = make([]int, len(arr3))
	copy3 = make([]int, len(arr3))
	copy(copy1, arr3)
	copy(copy2, arr3)
	copy(copy3, arr3)

	go InsertionSort(copy1, insertionSortChanArr)
	go SelectionSort(copy2, selectionSortChanArr)
	go BubbleSort(copy3, bubbleSortChanArr)

	fmt.Println()
	fmt.Println("=================")
	fmt.Println("Resultados 'arr3'")
	fmt.Println("=================")
	fmt.Printf("Ordenamiento por inserción: %d microsegundos\n", <-insertionSortChanArr)
	fmt.Printf("Ordenamiento por selección: %d microsegundos\n", <-selectionSortChanArr)
	fmt.Printf("Ordenamiento por burbuja: %d microsegundos\n", <-bubbleSortChanArr)
}

func InsertionSort(arr []int, duration chan int64) {
	initTime := time.Now()
	for i := 1; i < len(arr); i++ {
		auxVal := arr[i]
		for j := i - 1; j >= 0 && arr[j] > auxVal; j-- {
			arr[j+1] = arr[j]
			arr[j] = auxVal
		}
	}
	endTime := time.Now()
	durationInMicroseconds := endTime.Sub(initTime).Microseconds()
	duration <- durationInMicroseconds
}

func SelectionSort(arr []int, duration chan int64) {
	initTime := time.Now()
	var lengthOfArr = len(arr)
	for i := 0; i < lengthOfArr; i++ {
		var minValueIndex = i
		for j := i; j < lengthOfArr; j++ {
			if arr[j] < arr[minValueIndex] {
				minValueIndex = j
			}
		}
		arr[i], arr[minValueIndex] = arr[minValueIndex], arr[i]
	}
	endTime := time.Now()
	durationInMicroseconds := endTime.Sub(initTime).Microseconds()
	duration <- durationInMicroseconds
}

func BubbleSort(arr []int, duration chan int64) {
	initTime := time.Now()
	var lengthOfArr = len(arr)
	for i := 0; i < lengthOfArr-1; i++ {
		for j := 0; j < lengthOfArr-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	endTime := time.Now()
	durationInMicroseconds := endTime.Sub(initTime).Microseconds()
	duration <- durationInMicroseconds
}
