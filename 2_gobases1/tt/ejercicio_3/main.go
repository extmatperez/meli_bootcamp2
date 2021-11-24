package main

import "fmt"

func main() {

	age := 27
	isEmployed := true
	employmentAntiquity := 2
	salary := 105000.0

	if age < 22 {
		fmt.Println("Solo se otorgan prestamos a clientes mayores a 22 años")

	} else if !isEmployed {
		fmt.Println("Solo se otorgan presamos a clientes con empleo")

	} else if employmentAntiquity < 1 {
		fmt.Println("Solo se entregan prestamos a clientes que tengan al menos 1 año de antiguedad en su trabajo")

	} else {
		fmt.Println("Usted es apto para un préstamo")
		if salary > 100000 {
			fmt.Println("Además, por su buen salario no se le cobrará intereses")
		}
	}

}
