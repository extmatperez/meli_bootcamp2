package main

import "fmt"

func imprimir_nombre_direccion(nombre string, direccion string) {
	fmt.Println("Nombre y direccion")
	fmt.Printf("Nombre: %s \nDireccion: %s\n", nombre, direccion)
}

func imprimir_clima() {
	var temperatura float32 = 35.6
	var humedad float32 = 85
	var presion int = 1013
	fmt.Println("Clima")
	fmt.Println("Temperatura:", temperatura, "ºC")
	fmt.Println("Humedad:", humedad, "%")
	fmt.Println("Presion:", presion, "Hpa")

}

// EJRCICIO 3
// var 1nombre string    	I  sacar el 1
//   var apellido string 	C
//   var int edad       	I   var edad int
//   1apellido := 6    		I    scar el 1 y que sea string
//   var licencia_de_conducir = true   	I  bool al final
//   var estatura de la persona int  	C
//   cantidadDeHijos := 2       		C

// EJERCICIO 4
// var apellido string = "Gomez"
// var edad int = 35
// var booleano boolean = false
// var sueldo float = 45857.90
// var nombre string = "Julián"
