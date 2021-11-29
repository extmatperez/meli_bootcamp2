package main

import (
	"fmt"
	"math/rand"
	"time"
)

//"math/rand"

func main() {

	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	var copies []int

	c := make(chan int)
	c1 := make(chan int)
	c2 := make(chan int)
	fmt.Print()

	fmt.Println("-------Insertion-------")
	ini := time.Now()
	copy(copies, variable1)
	go insertion(variable1, c)
	go insertion(variable2, c1)
	go insertion(variable3, c2)
	tiempo1 := <-c
	tiempo2 := <-c1
	tiempo3 := <-c2
	fmt.Println("Insertion time:", tiempo1+tiempo2+tiempo3)
	fin := time.Now()
	fmt.Println("Run time:", fin.Sub(ini))
	fmt.Print()

	fmt.Println("-------Bubble-------")
	ini = time.Now()
	go bubble(variable1, c)
	go bubble(variable2, c1)
	go bubble(variable3, c2)
	fmt.Println("Tiempo de bubble:", tiempo1+tiempo2+tiempo3)
	fin = time.Now()
	fmt.Println("Tiempo de ejecución:", fin.Sub(ini))
	fmt.Print()

	fmt.Println("-------Selection-------")
	ini = time.Now()
	go selection(variable1, c)
	go selection(variable2, c1)
	go selection(variable3, c2)
	fmt.Println("Selection time:", tiempo1+tiempo2+tiempo3)
	fin = time.Now()
	fmt.Println("Run time:", fin.Sub(ini))
}

func bubble(num []int, c chan int) {
	begin := time.Now()
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	time.Sleep(1000 * time.Millisecond)
	end := time.Now()
	total := int(end.Sub(begin).Milliseconds())

	c <- total
}

func insertion(num []int, c chan int) {
	begin := time.Now()
	for i := 0; i < len(num); i++ {
		for j := i; j > 0 && num[j-1] > num[j]; j-- {
			num[j], num[j-1] = num[j-1], num[j]
		}
	}
	time.Sleep(1000 * time.Millisecond)
	end := time.Now()
	total := int(end.Sub(begin).Milliseconds())

	c <- total
}

func selection(num []int, c chan int) {
	begin := time.Now()
	for i := 0; i < len(num)-1; i++ {
		min := i
		for j := i + 1; j < len(num); j++ {
			if num[j] < num[min] {
				num[j], num[min] = num[min], num[j]
			}
		}
	}
	time.Sleep(1000 * time.Millisecond)
	end := time.Now()
	total := int(end.Sub(begin).Milliseconds())

	c <- total

}
