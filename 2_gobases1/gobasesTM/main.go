package main

import "fmt"

func main() {
	//ejercicio 1

	var nombre string = "Rodrigo Romero"
	var direccion string = "Vergani 463"

	fmt.Println("mi nombre es: " + nombre + " y mi direccion es :" + direccion)

	// ejercicio 2

	var temperatura int = 24
	var humedad int = 43
	var presionAtmosferica int = 1019

	fmt.Println("la temperatura es: ", temperatura, " y humedad es :", humedad, "%  y la presion atmosferica es: ", presionAtmosferica, "hPa")

	//ejercio 3

	//var nombre string// no se puede empezar una variable con numeros
	//var apellido string
	// var edad int // el nombre de la var va antes que el tipo de datos
	//apellidoDos := 6 // idem que primer variable  pero aparte se asigno un numero cuando debe ser un string, si se corrige el nombre de la variable va a a haber otro problema porque el nombre ya fue utilizado
	//var licencia_de_conducir = true // le falta la declaracion del tipo de dato bool
	// var estatura de la persona  int // no se puede crear una variable cuyo nombre tenga espacios
	//cantidadDeHijos := 2 // OK

	//ejercicio 4

	//var apellido string ="Gomez"  // OK
	// var edad int = "35" // tiene que ser un tipo de dato integer 35, se le asigno un string
	//boolean := "false"; // debe ser boolean:= false    booleano no es string y no se usa el ;
	// var sueldo string = 45857.90 // esto debe ser un float64 el tipo de numero ya que tiene decimales
	// var nombre string = "Julian" // OK
}
