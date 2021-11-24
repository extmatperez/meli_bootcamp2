package main

import "fmt"

func main() {

	//.1
	//bad: var 1nombre string
	//correct form:
	var nombre string
	nombre = "Mari"
	fmt.Println("Nombre:", nombre)
	//2.
	//Good
	var apellido string
	apellido = "Petit"
	fmt.Println("Apellido:", apellido)
	//3.
	//bad: var int edad
	//correct form:
	var edad int
	edad = 20
	fmt.Println("Edad:", edad)
	//4.
	//bad: 1apellido := 6
	//correct form:
	apellidoUno := "Gonzalez"
	fmt.Println("Apellido uno:", apellidoUno)
	//5.
	//bad: var licencia_de_conducir = true
	//correct form:
	var licenciaDeConducir = true
	fmt.Println("Licencia de conducir:", licenciaDeConducir)
	//6.
	//bad: var estatura de la persona int
	//correct form:
	var estaturaDeLaPersona int
	fmt.Println("Estatura de la persona:", estaturaDeLaPersona)
	//7.
	//Good
	cantidadDeHijos := 2
	fmt.Println("Cantidad de hijos:", cantidadDeHijos)

}
