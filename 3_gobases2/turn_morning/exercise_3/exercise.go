package main

// Chocolate Company
import (
	"errors"
	"fmt"
)

func main() {

	var category byte
	fmt.Println("¿Cual es la categoria?: ")
	fmt.Scanf("%c\n", &category)
	var bSalary int
	fmt.Println("Minitos trabajados: ")
	fmt.Scanf("%d\n", &bSalary)

	salary, err := getSalary(category, bSalary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario segun sus horas trabajadas es: %.2f", salary)
	}

}

func getSalary(category byte, minutes int) (float64, error) {

	switch category {
	case 'A':
		return float64(minutes * ((3000 / 60) * 1.5)), nil
	case 'B':
		return float64(minutes * ((1500 / 60) * 1.2)), nil
	case 'C':
		return float64(minutes * (1000 / 60)), nil
	default:
		return 0.0, errors.New("La categoria no es la correcta")

	}

}
