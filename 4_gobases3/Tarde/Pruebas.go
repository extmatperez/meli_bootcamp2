package main

import (
	"fmt"
	"time"
)

func proceso(i int, c chan int) {
	fmt.Println(i, "- Inicia")
	time.Sleep(10000 * time.Millisecond)
	fmt.Println(i, "- Termina")
	c <- time.Now().Second()
}

func main() {

	fmt.Println("PRUEBAS")

	c := make(chan int)
	go proceso(2, c)

	/*
		for i := 0; i < 10; i++ {
			go proceso(i)
		}
	*/

	variable := <-c
	fmt.Println("Termino el programa en", variable)
	fmt.Println("Variable", variable)
}
