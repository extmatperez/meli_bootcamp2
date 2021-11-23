package main

import "fmt"

func main() {
	var unNombre string                  //var 1nombre string //incorrecta
	var apellido string                  // correcta
	var edad int                         //var int edad //incorrecta
	unApellido := "apellido"             //1apellido := 6 //incorrecta
	var licencia_de_conducir bool = true //var licencia_de_conducir = true //incorrecta
	var estatura_de_la_persona int       //var estatura de la persona int //incorrecta
	cantidadDeHijos := 2                 //correcta

	fmt.Println(unNombre, apellido, edad, unApellido, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)
}
