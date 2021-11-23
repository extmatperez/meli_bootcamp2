package main

import "fmt"

func main() {
	var meses = [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto",
		"Septiembre", "Octubre", "Noviembre", "Diciembre"}

	var numeroMes int = 7
	fmt.Println(meses[numeroMes-1])
}

/*
	2. Es posible solucionarlo de varias formas:
		- Con un map donde la clave sea el número del mes y el valor el nombre.
		- Por medio de un switch.
		- Array declarado con los 12 meses y restando número de mes para que coincida con las posiciones del arreglo.
	Escogí la tercera forma ya que era la que menos código requería.
*/
