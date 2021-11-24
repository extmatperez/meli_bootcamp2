package main

import "fmt"

func impuesto(salario float64) (float64, string) {
	if salario <= 50.000 {
		return salario, "No se realizo ningun descuento"
	} else if salario > 50.000 && salario < 150.000 {
		return salario - salario*0.17, "Se desconto el 17%"
	} else {
		return salario - salario*0.27, "Se desconto el 27%"
	}
}

func main() {
	var salario float64
	fmt.Print("Cual es el salario: ")
	fmt.Scanf("%f", &salario)
	pago, info := impuesto(salario)
	fmt.Printf("El total a pagar es: %f -- %s\n", pago, info)
}
