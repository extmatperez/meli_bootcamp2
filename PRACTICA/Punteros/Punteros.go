package main

import (
	"fmt"
	"time"
)

func incrementar(dato *int) {
	*dato++
}

// ! sin canal
// func proceso(i int) {
// 	fmt.Println(i, "--Inicia")
// 	time.Sleep(1000 * time.Millisecond)
// 	fmt.Println(i, "--Termina")
// }

// ! Con canales
func procesoCanal(i int, c chan int) {
	fmt.Println(i, "--Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "--Termina")
	c <- i //? Enviamos el valor del canal
}

func main() {
	var a int = 25
	// fmt.Println(a)
	fmt.Printf("%v %d\n", a, a)

	// ? Declaramos un puntero b de tipo entero y le asignamos el valor de a
	var b *int = &a
	// fmt.Println(*b)
	fmt.Printf("Direccion en memoria: %v\nValor: %d\n", b, *b)

	incrementar(&a)
	fmt.Printf("%v %d\n", a, a)

	fmt.Println("*****************")

	// ! GO RUTINESS

	// for i := 0; i < 10; i++ {
	// 	go proceso(i)
	// }

	// time.Sleep(5000 * time.Millisecond)
	// fmt.Println("Termino el programa")

	// ! CANALES

	c := make(chan int)
	for i := 0; i < 10; i++ {
		go procesoCanal(1, c)
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Termino el programa en ", <-c)
	}
	// <-c
}
