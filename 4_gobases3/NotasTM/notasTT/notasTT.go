package main

import (
	"fmt"
	"time"
)

func main() {
	var p1 *int
	var num int

	//var p2 = new(*int)

	//fmt.Printf("Antes de inicializar... %v %v", p2, *p2)

	p1 = &num
	//fmt.Printf("\nAntes de inicializar... %v %v", p1, *p1)
	//fmt.Printf("\nAntes de inicializar... %v %v", &num, num)

	*p1 = 12
	//fmt.Printf("\nAntes de inicializar... %v %v", p1, *p1)
	//fmt.Printf("\nAntes de inicializar... %v %v", &num, num)

	fmt.Println()
	fmt.Println("Gorutine")
	fmt.Println()

	/*ini := time.Now()
	proceso(1)
	proceso(2)
	proceso(3)
	fin := time.Now()
	timepo := fin.Sub(ini)

	fmt.Println("El timepo secuencial que demoro es de: ", timepo.Seconds())*/

	ini := time.Now()
	for i := 0; i < 3; i++ {
		go proceso(i)
	}
	fin := time.Now()
	timepo := fin.Sub(ini)

	fmt.Println("El timepo paralelo que demoro es de: ", timepo.Seconds())
}

func proceso(i int) {
	fmt.Println(i, "inicia")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println(i, "termina")

}
