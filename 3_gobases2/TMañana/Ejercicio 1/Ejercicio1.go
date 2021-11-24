package main

import "fmt"

func impuesto(salario float64)(float64){

	if salario > 50000 && salario < 150000{
		return salario - (salario / 100 * 17.0)
	}else if salario >= 150000{
		return salario - (salario / 100 * 27.0)
	}
	return salario
}

func main(){

	var salario float64
	fmt.Println("Ingrese su salario base")
	fmt.Scanln(&salario)

	fmt.Println("Tu salario de bolsillo es:",impuesto(salario))
}