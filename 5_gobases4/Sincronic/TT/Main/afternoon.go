package main

import "fmt"

func main() {

	//Deffer
	func() {
		fmt.Println("Hola soy una funcion anonima")
	}() // Al agregar los parentecis al final se ejecuta inmediatamente esta funcion

	num := 15
	if num == 15 {
		panic("Panic Generado!")
	}

	func() {
		fmt.Println("Hola soy la segunda funcion anonima y no me ejecuto por que hay un panic anterior!")
	}()

}
