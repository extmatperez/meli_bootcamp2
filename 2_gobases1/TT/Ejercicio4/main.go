package main

import "fmt"

func main() {
	var meses = map[int]string{
		1:"Enero",
		2:"Febrero",
		3:"Marzo",
		4:"Abril",
		5:"Mayo",
		6:"Junio",
		7:"Julio",
		8:"Agosto",
		9:"Septiembre",
		10:"Octubre",
		11:"Noviembre",
		12:"Diciembre",
	}

	mes := 5

	fmt.Println("El mes correspondiente al número ingresado es: ", meses[mes])
}