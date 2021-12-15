package main

/*


import "fmt"

func Incrementar(c *int) {
	*c++

}

func main() {
	// var v int = 19
	// var p *int

	// p = &v
	// fmt.Printf("El puntero p referencia a la dirección de memoria: %v \n", p)
	// fmt.Printf("Al desreferenciar el puntero p obtengo el valor: %d \n", *p)

	var c int = 19

	Incrementar(&c)
	fmt.Println("El valor de v ahora vale:", c)
}
*/

// func main() {
// 	var p1 *int
// 	var num int
// 	num = 5

// 	fmt.Printf("Antes de inicializar %v, %v\n", p1, num)

// 	p1 = &num

// 	fmt.Printf("Despues e inicializar %v, %v\n", p1, *p1)
// 	fmt.Printf("Despues de inicializar %v, %v\n", &num, num)
// }

// import (
// 	"fmt"
// 	"time"
// )

// func proceso(i int) {
// 	fmt.Println(i, "-inicia")
// 	time.Sleep(1000 * time.Millisecond)
// 	fmt.Println(i, "/termina")
// }

// func main() {
// 	ini := time.Now()
// 	for i := 0; i < 3; i++ {
// 		go proceso(i)
// 	}

// 	fin := time.Now()

// 	tiempo := fin.Sub(ini)

// 	fmt.Println("El tiempo demorado es de ", tiempo.Seconds())
// }

// //canales

//  func proceso(i int, c chan int) {
// 	fmt.Println(i, "-inicia")
// 	time.Sleep(1000 * time.Millisecond)
// 	fmt.Println(i, "/termina")
// 	c <- i
// }

// func main() {

// 	c := make(chan int)

// 	ini := time.Now()
// 	for i := 0; i < 10; i++ {
// 		go proceso(i, c)

// 	}
// 	variable := <-c
// 	fin := time.Now()

// 	tiempo := fin.Sub(ini)
// 	fmt.Println(variable)
// 	fmt.Println("El tiempo paralelo demorado es de :", tiempo.Seconds())
// }

/* import (
	"fmt"
	"time"
)

func proceso(i int) {
	fmt.Println(i, "/inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-termina")
}

func main() {

	for i := 0; i < 10; i++ {
		proceso(i)
	}
}
*/

import (
	"fmt"
	"time"
)

func proceso(i int) {
	fmt.Println(i, "/inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-termina")
}

func main() {

	for i := 0; i < 10; i++ {
		go proceso(i)
	}
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("-termino el programa")
}
