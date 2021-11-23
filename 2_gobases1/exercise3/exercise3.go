// Declaracion de variables

/*
Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles las correspondientes devoluciones.
Uno de los puntos del examen consiste en declarar distintas variables.

Necesita ayuda para:

Detectar cuáles de estas variables que declaró el alumno son correctas:
- apellido
- cantidadDeHijos


Corregir las incorrectas:

var 1nombre string 					-> var nombre string
var apellido string					->
var int edad						-> var edad int
1apellido := 6						-> primerApellido := 6 entendiendose que desea una nueva variable numerica
var licencia_de_conducir = true		-> var licencia_de_conducir bool = true
var estatura de la persona int		-> var estatura_de_la_persona int
cantidadDeHijos := 2				->



*/
package main

import "fmt"

func main() {
	var nombre string
	var apellido string
	var edad int
	var licencia_de_conducir bool = true
	var estatura_de_la_persona int
	cantidadDeHijos := 2

	nombre = "jose"
	apellido = "rios"
	edad = 28
	estatura_de_la_persona = 170
	cantidadDeHijos = 0
	fmt.Printf("nombre completo: %s %s, edad: %d, licencia de conducir? %v, estatura: %d, cantidad de hijos: %d", nombre, apellido, edad, licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)
}
