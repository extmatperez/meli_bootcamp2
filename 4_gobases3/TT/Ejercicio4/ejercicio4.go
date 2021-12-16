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

	c := make(chan int64)

	go insertion_sort(&variable1, c)
	fmt.Printf("%d\n", (c))
	go insertion_sort(&variable1, c)
	fmt.Printf("%d\n", c)
	go insertion_sort(&variable1, c)
	fmt.Printf("%d\n", c)

	go insertion_sort(&variable2, c)
	fmt.Printf("%d\n", c)
	go insertion_sort(&variable2, c)
	fmt.Printf("%d\n", c)
	go insertion_sort(&variable2, c)
	fmt.Printf("%d\n", c)

	go insertion_sort(&variable3, c)
	fmt.Printf("%d\n", c)
	go insertion_sort(&variable3, c)
	fmt.Printf("%d\n", c)
	go insertion_sort(&variable3, c)
	fmt.Printf("%d\n", c)
}

func buble_sort(numeros *[]int, c chan int64) {
	inicio := time.Now()
	for i := len(*numeros); i > 0; i-- {
		for j := 1; j < i; j++ {
			if (*numeros)[j-1] > (*numeros)[j] {
				intermediate := (*numeros)[j]
				(*numeros)[j] = (*numeros)[j-1]
				(*numeros)[j-1] = intermediate
			}
		}
	}
	fin := time.Now()
	delta := fin.Sub(inicio)
	c <- delta.Microseconds()
}

func insertion_sort(numeros *[]int, c chan int64) {
	inicio := time.Now()
	var n = len(*numeros)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if (*numeros)[j-1] > (*numeros)[j] {
				(*numeros)[j-1], (*numeros)[j] = (*numeros)[j], (*numeros)[j-1]
			}
			j = j - 1
		}
	}
	fin := time.Now()
	delta := fin.Sub(inicio)
	c <- delta.Microseconds()
}

func selection_sort(numeros *[]int, c chan int64) {
	inicio := time.Now()
	var n = len((*numeros))
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if (*numeros)[j] < (*numeros)[minIdx] {
				minIdx = j
			}
		}
		(*numeros)[i], (*numeros)[minIdx] = (*numeros)[minIdx], (*numeros)[i]
	}
	fin := time.Now()
	delta := fin.Sub(inicio)
	c <- delta.Microseconds()
}
