package main

import "fmt"

func main() {
	var month int
	months := []string{"", "Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	fmt.Printf("%v,%s", month, months[month])
}
