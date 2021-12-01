/*
Realizar una aplicación que contenga una variable con el número del mes.
Según el número, imprimir el mes que corresponda en texto.
¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
*/

package main

import "fmt"

func main() {

	numeroMes := 6

	meses := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	if 1 > numeroMes || numeroMes > 12{
		fmt.Println("El índice del mes no coincide con ningún mes")
	} else {
		fmt.Println(meses[numeroMes - 1])
	}
}