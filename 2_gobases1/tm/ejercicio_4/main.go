package main

import "fmt"

func main() {

	var apellido string = "Gomez"

	//"35" -> 35 (se sacan comillas, para convertir valor de string a int)
	var edad int = 35

	//tecnicamente correcto, pero por contexto asumo que quiere declarar
	//una variable de tipo boolean y no de tipo string, se quitan comillas
	boolean := false

	//se cambia tipo de variable, string -> float32
	var sueldo float32 = 45857.90

	var nombre string = "Julián"

	//agrego println para usar las variables y que el compilador no chille
	fmt.Println(apellido, edad, boolean, sueldo, nombre)
}
