package main

import (
	"fmt"
	"time"
)

func proceso(i int, c chan int) {
	fmt.Println(i, "-inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-termina")
	c <- i
}

func main() {
	c := make(chan int)
	suma := 0
	//Forma Thread:
	iniA := time.Now()
	for i := 0; i < 10; i++ {
		go proceso(i, c)
	}
	for i := 0; i < 10; i++ {
		variable := <-c
		suma += variable
		fmt.Println("Termino el programa en ", variable)

	}
	finA := time.Now()
	tiempoA := finA.Sub(iniA)

	fmt.Println("El tiempo demorado de forma paralela es: ", tiempoA.Seconds())
}
