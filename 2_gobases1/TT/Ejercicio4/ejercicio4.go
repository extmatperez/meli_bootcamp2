package main

import "fmt"

func main() {
	month := 11

	switch month {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("Mes inexistente")
	}

	// Otra forma que se me ocurre es con un array y es la que eligiria
	monthsArray := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	if month >= 1 && month <= 12 {
		fmt.Println(monthsArray[month-1])
	} else {
		fmt.Println("Mes inexistente")
	}

	// Otra forma que se me ocurre es con un mapa
	monthsMap := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	monthName, exists := monthsMap[month]
	if exists {
		fmt.Println(monthName)
	} else {
		fmt.Println("Mes inexistente")
	}
}
