package main

// Chocolate Company
import (
	"fmt"
)

func main() {
	var salary float64
	fmt.Println("Please enter your salary: ")
	fmt.Scanf("%f.2\n", &salary)

	fmt.Println(getTax(salary))
}

func getTax(salary float64) string {
	var tax float64
	if salary > 50000 {
		tax = salary * .17
		return "La persona tiene un impuesto de " + fmt.Sprintf("%.2f", tax)
	} else if salary > 150000 {
		tax = salary * .27
		return "La persona tiene un impuesto de " + fmt.Sprintf("%.2f", tax)
	} else {
		return "La persona no tendra impuestos: " + fmt.Sprintf("%.2f", tax)
	}

}
