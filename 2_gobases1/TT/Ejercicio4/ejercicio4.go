package main

import "fmt"

func main() {
	fmt.Println(month(1))
	fmt.Println(month(3))
	fmt.Println(month(12))
	fmt.Println(month(10))
	fmt.Println(month(7))
	fmt.Println(month(-1))
}

func month(n int) string {
	switch n {
	case 1:
		return "Enero"
	case 2:
		return "Febrero"
	case 3:
		return "Marzo"
	case 4:
		return "Abril"
	case 5:
		return "Mayo"
	case 6:
		return "Junio"
	case 7:
		return "Julio"
	case 8:
		return "Agosto"
	case 9:
		return "Setiembre"
	case 10:
		return "Octubre"
	case 11:
		return "Noviembre"
	case 12:
		return "Diciembre"
	}
	return "Seleccione un numero valido para obtener un mes"
}
