package main

import (
	"errors"
	"fmt"
)


func salaryCalc(salary int)error{
	if salary<150000{
		return errors.New("100 - No alcanza el minimo imponible")
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