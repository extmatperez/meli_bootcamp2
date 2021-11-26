package main

import (
	"fmt"
)


func ProcesSalary(salary int) (error) {

	if salary < 150000 {
		return fmt.Errorf("400 - el mínimo imponible es de 150.000 y el salario ingresado es de: %v" ,salary)
		}
	return nil
}

func PagaImpuesto (salary int){
	err := ProcesSalary(salary) 

	if(err != nil) {
		fmt.Println(err)
	}else{

		fmt.Println("Debe pagar impuesto")
	}
}



func main() {

	PagaImpuesto(10)
	PagaImpuesto(210000)
 }


