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
	ini := time.Now()
	proceso(1)
	proceso(2)
	proceso(3)
	proceso(4)
	fin := time.Now()
	tiempo := fin.Sub(ini)

	fmt.Println("El tiempo demorado es de: ", tiempo.Seconds())
}
