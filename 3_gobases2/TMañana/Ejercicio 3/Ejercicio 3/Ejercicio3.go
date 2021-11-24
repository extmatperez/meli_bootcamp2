package main

import "fmt"

func min_a_h(minutos int) int {
	return minutos / 60
}

func salario_mensual(base int, horas int) int {
	return (base * horas)
}

func calc_salario(cat string, minutos int) (int){
	horas := min_a_h(minutos)
	switch cat {
		case "A", "a":
			return (salario_mensual(3000, horas)) + ((salario_mensual(3000, horas))/100*50)
		case "B", "b":
			return (salario_mensual(1500, horas)) + ((salario_mensual(1500, horas))/100*20)
		default:
			return (1000*min_a_h(minutos))
		
	}
}

func main() {
	fmt.Println("Su salario es: ", calc_salario("A", 5400))
	fmt.Println("Su salario es: ", calc_salario("B", 3200))
	fmt.Println("Su salario es: ", calc_salario("C", 4200))
}