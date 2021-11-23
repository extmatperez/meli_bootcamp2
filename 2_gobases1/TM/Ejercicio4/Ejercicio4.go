package main

import "fmt"

func main() {
	var apellido string = "Gomez" // Este esta bien.
	//var edad int = "35" // Esto esta mal, los enteros van sin comillas
	var edad int = 35
	//boolean := "false"; // Esto esta mal, true o false van sin comillas, aparte sin puntocoma al final, no es admitido por Go.
	// En realidad esta bien porque seria un string llamado boolean con valor "false", no como booleano sino como valor de cadena.
	boolean := false
	//var sueldo string = 45857.90 // Esto esta mal, se trata de un flotante
	var sueldo float64 = 45857.90
	var nombre string = "Julián"

	fmt.Println(apellido, edad, boolean, sueldo, nombre)

}
