package main

import (
	"errors"
	"fmt"
	"sort"
)

type MyError struct{

	status int
	msg string
}

func (e *MyError) Error() string {

	return fmt.Sprintf("%d - %v",e.status,e.msg)
}


func HaveMininHours(hoursWorked int) (error) {

	if hoursWorked <= 80 || hoursWorked < 0  {
		return &MyError{
			status: 400,
			msg: "Error en HaveMininHours el trabajador no puede haber trabajado menos de 80 hs mensuales",
		}
		}	
	return nil
}

func SalaryDiscount(salary float64) (float64){
	
	if(salary >= 150000){
		return salary * (1-0.1)
	}
	return salary
}

func ProccesSalary (hoursWorked int, priceHour float64)  (float64,error){
	err1 := HaveMininHours(hoursWorked)
	if(err1 != nil){
		err2 := fmt.Errorf("ErrorPadre: la cantidad de horas %v no es suficiente. \nErrorHijo: %w",hoursWorked,err1)
		errorUnwraped := errors.Unwrap(err2)
		return 0, fmt.Errorf("\nErrorPadre: %w , \nErrorUnwrapped: %v",err2,errorUnwraped)
	}
	salary := SalaryDiscount(float64(hoursWorked)*priceHour)

	return salary,nil
}

func getAguinaldo (salaries ... float64) (float64,error){

	numSalaries :=  len(salaries)
	if(numSalaries == 0 ) {
		return 0, errors.New("debe tener al menos un salario para calcular el aguinaldo")
	}
	sort.Float64s(salaries)
	salaryMax := salaries[0]
	return salaryMax/ float64((12 * numSalaries)),nil 


}


func main() {

	salario,err := ProccesSalary(70,1500)

	if(err != nil){
		fmt.Print(err)

	}else{
		fmt.Printf("El salario es: %v", salario)
	}



	//////////////////////////////////////////////////


	aguinaldo,err := getAguinaldo()

	if( err != nil) {
		fmt.Println("\nError en aguinaldo: ",err)
	}else {
	fmt.Printf("\nEl aguinaldo es: %v y el medio aguinadlo es: %v", aguinaldo, aguinaldo/2)
	}



 }


