package main

import "fmt"

// Meses
func main() {

	month := 10

	//Approach 1
	
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
		fmt.Println("Mes Desconocido")
	}

	//Approach 2
	month_map := map[string]int{"Enero":1, "Febrero":2, "Marzo":3, "Abril":4,"Mayo":5,"Junio":6,"Julio":7,"Agosto":8,"Septiembre":9,"Octubre":10,"Noviembre":11,"Diciembre":12}

	for k, v := range month_map {
		if month == v {
			fmt.Println("Mes:", k)
			break
		}
	}
}
