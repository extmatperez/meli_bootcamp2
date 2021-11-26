package main

import (
	"fmt"
)


func salaryCalc(salary int)error{
	if salary<150000{
		return fmt.Errorf("El minimo imponible es de $150000 y el salario ingresado es de: $%v", salary)
	}else{
		return nil

	}

}

func main(){
	var salary int
	fmt.Scanln(&salary)
	err := salaryCalc(salary)

	if err != nil {
		fmt.Printf("Ha ocurrido un error: %v\n", err)
	}else{
		fmt.Printf("Debe pagar impuestos")
	}

}