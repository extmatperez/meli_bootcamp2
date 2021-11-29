// PANIC y para que sirven.
// Panic es una interrupcion de la ejecucion del programa, con la indicacion de que algo salio mal de forma inesperada. Se corta donde ocurrio el error.
// Por ejemplo, con un out of bounds en un array, cuando invocamos metodos en punteros nulos o intentar trabajar con canales cerrados.
// Osea basicamente son errores que el compilador no puede captar en tiempo de ejecucion. Rompe la ejecucion de la aplicacion, por supuesto.
// PANIC es lo mas usado para realizar debug.

// Hay dos sentencias incorporadas, DEFER y RECOVER. Sirven para controlar los efectos de un panic y evitar que el programa finalice de modo no deseado.
// Si bien son independientes, hace falta que se usen ambas complementariamente para lograr resultados de mejor performance.

// DEFER
// Nos permite diferir la ejecucion de ciertas funciones y asegurar que sean ejecutadas antes de la finalizacion de la ejecucion del programa.
// Nos asegura que ciertas funciones seran ejecutadas a pesar de que ocurra un panic. Por ejemplo, cerrar una conexion a la base de datos.
// Nos aseguramos de limpiar recursos durante la ejecucion del programa. Se usa como mecanismo de seguridad.
// Los defer NO se ejecutan ante un log.Fatal()

// RECOVER
// Si ocurrio en algun lado un panic, vamos a guardar la descripcion del panic en un error.

package main

import "fmt"

type Dog struct {
	name string //
}

func (s *Dog) woofWoof() {
	fmt.Println(s.name, " hace woof woof!")
}

func main() {
	// animals := []string{"vaca", "perro", "halcon"}
	// fmt.Println(animals[3]) // Aqui ocurre una excepcion tipo panic! Porque desea acceder a un indice inexistente en ese slice.
	// s := &Dog{"Sammy"}
	// s = nil
	// s.woofWoof() // Aqui ocurre otro panic, porque desea llamar un metodo de una estructura que esta seteada en nil.

	// Estos ultimos dos casos de panic son los ejemplos mas comunes de errores en tiempos de ejecucion que tiran un panic.

	// Ejemplo de DEFER antes de panic que si se va a mostrar.
	defer func() {
		fmt.Println("Esta linea se va a imprimir a pesar de que ocurra un panic 1.")
	}()
	defer func() {
		fmt.Println("Esta linea se va a imprimir a pesar de que ocurra un panic 2.")
	}()
	defer func() {
		fmt.Println("Esta linea se va a imprimir a pesar de que ocurra un panic 3.")
	}()
	func() {
		fmt.Println("Esta linea se va a imprimir a pesar de que ocurra un panic 4.")
	}()
	func() {
		fmt.Println("Esta linea se va a imprimir a pesar de que ocurra un panic 5.")
	}()
	func() {
		fmt.Println("Esta linea se va a imprimir a pesar de que ocurra un panic 6.")
	}()
	// En este ejemplo vemos que se forma una pila LIFO (ultimo entrado primero en salir), por lo que si se encuentra un panic
	// los DEFER se ejecutan en orden inverso en el que se definieron.
	// Los que no son difer, que son solo funciones anonimas, se ejecutan primero y en el mismo orden en el que se definieron.

	// Defer es un ejemplo de funcion anonima. No necesita nombre y se corre si o si con el solo hecho de definirla.
	panic("Ocurre un panic!")

	// Ejemplo de DEFER que NO se va a mostrar porque viene despues de un panic.
	/*defer func() {
		fmt.Println("Esta linea se va a imprimir a pesar de que ocurra un panic.")
	}()*/
}
