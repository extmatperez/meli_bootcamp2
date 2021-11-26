package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	ini := time.Now()

	for i := 0; i < 10; i++ {
		go proceso(i, c)
	}

	for i := 0; i < 10; i++ {
		//variable := <-c
		//fmt.Println(variable)
		fmt.Println("Termino el programa en ", <-c)

	}
	fin := time.Now()
	timepo := fin.Sub(ini)

	fmt.Println("El timepo paralelo que demoro es de: ", timepo.Seconds())

}

func proceso(i int, c chan int) {
	fmt.Println(i, "inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "termina")
	c <- i
}
