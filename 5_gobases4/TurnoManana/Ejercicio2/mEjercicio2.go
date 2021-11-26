package main

import (
	"errors"
	"fmt"
)


func ProcesSalary(salary int) (error) {

	if salary < 150000 {
		return errors.New("400 - el salario ingresado no alcanza el mínimo imponible")
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


