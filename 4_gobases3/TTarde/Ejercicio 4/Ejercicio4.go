package main

import (
	"fmt"
	"math/rand"
	"time"
)

func burbuja(num []int, c chan int) {
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

func insercion(num []int, c chan int) {
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

func seleccion(num []int, c chan int) {
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

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)
	var copia []int

	c := make(chan int)
	c1 := make(chan int)
	c2 := make(chan int)
    ini := time.Now()
	copy(copia, variable1)
	go burbuja(variable1, c)
	go burbuja(variable2, c1)
	go burbuja(variable3, c2)
    tiempo1 := <-c
    tiempo2 := <-c1
	tiempo3 := <-c2
	fmt.Println("Tiempo de burbuja:", tiempo1+tiempo2+tiempo3)
    //fmt.Println("Tiempo de burbuja:", <-c+<-c1+<-c2)
    fin := time.Now()
    fmt.Println("Tiempo de ejecución:", fin.Sub(ini))

    
    ini = time.Now()
	go seleccion(variable1, c)
	go seleccion(variable2, c1)
	go seleccion(variable3, c2)
	fmt.Println("Tiempo de seleccion:", tiempo1+tiempo2+tiempo3)
    tiempo1 = <-c
    tiempo2 = <-c1
	tiempo3 = <-c2
	fmt.Println("Tiempo de burbuja:", tiempo1+tiempo2+tiempo3)
    fin = time.Now()
    fmt.Println("Tiempo de ejecución:", fin.Sub(ini))

    
    ini = time.Now()
	go insercion(variable1, c)
	go insercion(variable2, c1)
	go insercion(variable3, c2)
    tiempo1 = <-c
    tiempo2 = <-c1
	tiempo3 = <-c2
	fmt.Println("Tiempo de burbuja:", tiempo1+tiempo2+tiempo3)
    fin = time.Now()
    fmt.Println("Tiempo de ejecución:", fin.Sub(ini))
	fmt.Println("Tiempo de insercion:", tiempo1+tiempo2+tiempo3)

}
