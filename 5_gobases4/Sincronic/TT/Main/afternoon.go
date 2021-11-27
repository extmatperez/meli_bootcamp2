package main

import "fmt"

func main() {
	var num int = 0
	//Deffer
	func() {
		fmt.Println("1 Hola soy una funcion anonima")
	}() // Al agregar los parentecis al final se ejecuta inmediatamente esta funcion

	defer func() {
		fmt.Println("2 Hola soy la segunda funcion anonima y no me ejecuto por que hay un panic anterior!")
	}()

	defer func() {
		fmt.Println("3 Hola soy la tercera", num)
	}()

	defer func() {
		fmt.Println("4 Hola soy la tercera funcion anonima", num)
	}()

	func() {
		fmt.Println("5 Hola soy una funcion anonima", num)
	}()

	num = 10
	panic("Panic Generado!")

}
