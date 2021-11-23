package main

import "fmt"

func main() {

	var apellido string = "Gomez" //correcta
	fmt.Println(apellido)

	//var edad int = "35"; -- Esta asignando un valor erroneo a un entero, debe ir sin comillas y en go no se usa el ;
	var edad int = 35
	fmt.Println(edad)

	//boolean := "false" -- Los valores booleanos se asignan sin comillas
	boolean := false
	fmt.Println(boolean)

	//var sueldo string = 45857.50 esta asignando un valor numerico con decimales, deberia ponerle las comillas o usar un tipo float como dato
	var sueldo float64 = 45857.50
	fmt.Println(sueldo)

	var nombre string = "Julian" //correcta
	fmt.Println(nombre)
}
