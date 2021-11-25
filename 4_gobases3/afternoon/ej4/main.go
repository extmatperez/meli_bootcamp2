package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Times struct {
	name       string
	passedTime int64
}

func sortByInsertion(numbers []int, c chan Times) {
	init := time.Now()

	for i := 1; i < len(numbers); i++ {
		actual := numbers[i]
		j := i - 1
		for (j >= 0) && (numbers[j] > actual) {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = actual
	}
	end := time.Now()
	diff := end.Sub(init).Microseconds()
	c <- Times{"Insertion", diff}
}

func sortBySelection(numbers []int, c chan Times) {
	init := time.Now()

	for i := 0; i < len(numbers)-1; i++ {
		p := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[p] {
				p = j
			}
		}

		numbers[p], numbers[i] = numbers[i], numbers[p]
	}
	end := time.Now()
	diff := end.Sub(init).Microseconds()
	c <- Times{"Selection", diff}
}

func bubbleSort(numbers []int, c chan Times) {
	init := time.Now()

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	end := time.Now()
	diff := end.Sub(init).Microseconds()
	c <- Times{"Bubble", diff}
}

func testSorts(n int, orders []string, c chan Times) {
	firstArray := rand.Perm(n)
	secondArray := rand.Perm(n)
	thirdArray := rand.Perm(n)

	go sortByInsertion(firstArray, c)
	go sortBySelection(secondArray, c)
	go bubbleSort(thirdArray, c)

	for i := 0; i < 3; i++ {
		data := <-c
		fmt.Printf("%v: %v en %d microsegundos\n", orders[i], data.name, data.passedTime)
	}
}

func main() {
	c := make(chan Times)
	orders := []string{"Primero", "Segundo", "Tercero"}

	testSorts(100, orders, c)
	testSorts(1000, orders, c)
	testSorts(10000, orders, c)
}
