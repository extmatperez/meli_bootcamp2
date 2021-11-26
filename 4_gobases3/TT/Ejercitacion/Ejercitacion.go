package main

// Uso de punteros: es un tipo de datos que tienen como valor una direccion a memoria que se refiere a otro valor almacenado.
// & es la direccion de esa variables
// * es para indicar un tipo de dato puntero.
// Entonces a ese puntero creado para un determinado dato se le puede asignar la direccion de una determinada variable.

import (
	"fmt"
	"runtime"
	"time"
)

func proceso(i int, c chan int) {
	fmt.Println("\n-Inicia", i)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "- termina")
	c <- i
}

func main() {
	var p1 *int
	var num int

	var p2 = new(int)
	*p2 = 65

	fmt.Printf("p1 - Antes de inicializar: %v %v\n", p1, num)
	fmt.Printf("\np2 - Antes de inicializar: %v %v", p2, num)

	p1 = &num
	num = 5
	fmt.Printf("\nDespues de inicializar: %v %v", p1, *p1)
	fmt.Printf("\nDespues de inicializar: %v %v\n", &num, num)

	*p1 = 15
	fmt.Printf("\nDespues de inicializar: %v %v", p1, *p1)
	fmt.Printf("\nDespues de inicializar: %v %v", &num, num)

	// GO ROUTINES
	/*proceso(1)
	ini := time.Now()
	proceso(10)
	proceso(10)
	proceso(15)
	proceso(10)
	fin := time.Now()
	fmt.Printf("\nTiempo de demora de proceso: %v", fin.Sub(ini))*/

	// Si queremos que sea tiempo paralelo, entonces es de la siguiente forma.
	// Vemos en la salida que los mas pequeños son los que se terminan mas rapido y los mas pesados al ultimo. Salen en orden.
	/*ini := time.Now()
	go proceso(1800000)
	go proceso(1000000)
	go proceso(1500000)
	go proceso(1200000)
	fin := time.Now()
	fmt.Printf("\nTiempo paralelo de demora de proceso: %v", fin.Sub(ini))
	*/
	fmt.Println("\nEl numero de CPU es ", runtime.NumCPU())

	// Despues tenemos los canales, que nos van a permitir enviar valores a las rutinas y esperar hasta recibir un valor.
	// CHANNELS
	c := make(chan int)

	ini := time.Now()
	go proceso(1, c)
	go proceso(2, c)
	go proceso(3, c)
	go proceso(4, c)
	go proceso(5, c)
	go proceso(6, c)
	go proceso(7, c)
	go proceso(8, c)
	go proceso(9, c)
	go proceso(10, c)

	variable := <-c
	fin := time.Now()
	fmt.Println("\nFinalizo el canal ", variable)
	fmt.Printf("\nTiempo paralelo de demora de proceso: %v", fin.Sub(ini).Seconds())

}
