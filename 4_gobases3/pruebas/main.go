package main

import (
	"fmt"
	"time"
)

func proceso(nro int, c chan string) {
	var valor string
	fmt.Printf("%d Ingrese un texto: \n", nro)
	fmt.Scanf("%s", &valor)
	fmt.Printf("%d Ingresado.\n", nro)
	c <- ""
}

func main() {

	c := make(chan string)

	for i := 0; i < 10; i++ {
		go proceso(i, c)
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		<-c
	}

}
