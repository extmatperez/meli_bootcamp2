package main

import "fmt"

func main() {
	var salary, taxes, aditionalTaxes, totalTaxes float64
	var answer string

	salary = 180000
	aditionalTaxes = 10
	taxes = 17 

	if salary > 50000{
		answer = "Impuestos del 17% aplicado al salario."
		totalTaxes = salary * (taxes/100)
	}
	if salary > 150000 {
		answer = "Impuestos aplicados: \n Impuesto común del 17% \nImpuesto adicional aplicado del 10%"
		totalTaxes += (aditionalTaxes/100)
	}

	fmt.Println(totalTaxes, "\n", answer)
}
