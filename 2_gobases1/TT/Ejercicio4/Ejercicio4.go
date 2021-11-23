package main

import (
	"fmt"
)

func main() {
	var month_num int
	fmt.Println("Elija un numero de mes: ")
	fmt.Scanf("%d", &month_num)
	switch month_num {
	case 1:
		fmt.Println(month_num, "Enero")
	case 2:
		fmt.Println(month_num, "Febrero")
	case 3:
		fmt.Println(month_num, "Marzo")
	case 4:
		fmt.Println(month_num, "Abril")
	case 5:
		fmt.Println(month_num, "Mayo")
	case 6:
		fmt.Println(month_num, "Junio")
	case 7:
		fmt.Println(month_num, "Julio")
	case 8:
		fmt.Println(month_num, "Agosto")
	case 9:
		fmt.Println(month_num, "Septiembre")
	case 10:
		fmt.Println(month_num, "Octubre")
	case 11:
		fmt.Println(month_num, "Noviembre")
	case 12:
		fmt.Println(month_num, "Diciembre")
	default:
		fmt.Println("El numero ingresado no corresponde a ningun mes.")
	}
}
