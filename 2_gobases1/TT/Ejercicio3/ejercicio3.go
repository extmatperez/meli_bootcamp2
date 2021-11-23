package main

import "fmt"

func main() {
	age := 22
	isEmployee := false
	antiquity := 365
	salary := 10000

	if age > 22 && isEmployee && antiquity > 365 {
		if salary > 100000 {
			fmt.Println("El préstamo se otorgara sin interes")
		} else {
			fmt.Println("El préstamo se otorgara con interes")
		}
	} else {
		if age <= 22 {
			fmt.Println("No se puede otorgar el préstamo a menores de 23 años")
		} else if !isEmployee {
			fmt.Println("No se puede otorgar el préstamo a desempleados")
		} else if antiquity <= 365 {
			fmt.Println("No se puede otorgar el préstamo a personas que no tengan mas de un año de antiguedad")
		}
	}
}
