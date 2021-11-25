package main

import (
	"fmt"
	"time"
)

func proceso(i int) {
	fmt.Println(i, "-inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-termina")
}

func main() {
	/*
		var num int = 6
		var p1 *int

		p1 = &num

		fmt.Printf("Punteros. Direccion: %v Valor del puntero: %v \n", p1, *p1)
	*/
	//*********************************************************************
	//************ GOROUTINES *********************************************
	//*********************************************************************

	//Forma Secuencial:
	iniS := time.Now()
	proceso(1)
	proceso(2)
	proceso(3)
	proceso(4)
	finS := time.Now()
	tiempoS := finS.Sub(iniS)

	fmt.Println("El tiempo demorado de forma secuencial es: ", tiempoS.Seconds())

	//Forma Thread:
	iniA := time.Now()
	go proceso(1)
	go proceso(2)
	go proceso(3)
	go proceso(4)
	finA := time.Now()
	tiempoA := finA.Sub(iniA)

	fmt.Println("El tiempo demorado de forma paralela es: ", tiempoA.Seconds())
}
