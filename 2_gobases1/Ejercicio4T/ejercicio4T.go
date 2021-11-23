package main

import "fmt"

func main() {

	mes := 11
	mesPalabra := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio",
		"Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	if mes > 1 && mes < 13 {
		fmt.Println("El mes", mes, "es", mesPalabra[mes-1])
	} else {
		fmt.Println("El mes", mes, "no existe")
	}
}
