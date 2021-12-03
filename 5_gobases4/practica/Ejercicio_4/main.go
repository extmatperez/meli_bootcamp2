package main

import (
	"errors"
	"fmt"
)

type Worker struct {
	name			string
	salaries		[]float64
}
func main() {
	/*
		Vamos a hacer que nuestro programa sea un poco más complejo y útil.
		Desarrolla las funciones necesarias para permitir a la empresa calcular:
		Salario mensual de un trabajador según la cantidad de horas trabajadas.
		La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
		Dicha función deberá retornar más de un valor (salario calculado y error).
		En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar
		el 10% en concepto de impuesto.
		En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
		la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber
		trabajado menos de 80 hs mensuales”.
		Calcular el medio aguinaldo correspondiente al trabajador
		Fórmula de cálculo de aguinaldo:
		[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
		La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso
		de que se ingrese un número negativo.

		Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando
		“errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los
		retornos de error en tu función “main()”.
	*/
	var (
		worker = Worker{"Lucia", nil}
		salary float64
		err error
	)
	fmt.Println("SALARIO CALCULATOR: ")
	fmt.Println("-------------------------------------------------------------------------------")
	salary, err = salaryCalculator(90, 1400, &worker)
	// for only to register salaries
	for i := 0; i<2; i++{
		salary, err = salaryCalculator(90, 1800, &worker)
	}

	salary, err = salaryCalculator(100, 1800, &worker)
	if err != nil {
		fmt.Println(err)
		fmt.Println("-------------------------------------------------------------------------------")

	} else {
		fmt.Println("Su ultimo salario registrado es: ", salary)
		fmt.Println("-------------------------------------------------------------------------------")
	}

	fbonus, errBonus := bonus(worker)
	fmt.Println("BONUS")
	fmt.Println("-------------------------------------------------------------------------------")
	if errBonus != nil {
		fmt.Println(errBonus)
		fmt.Println("-------------------------------------------------------------------------------")
	} else {
		fmt.Println("Su aguinaldo es: ", fbonus)
		fmt.Println("-------------------------------------------------------------------------------")
	}
}
func salaryCalculator(hours float64, hoursValue float64, worker *Worker) (float64, error) {
	var finalSalary float64= 0

	if hours <= 80 {
		return 0, errors.New("error: El trabajador no puede haber trabajado menos de 80 hrs mensuales")
	}
	finalSalary = hours * hoursValue
	if finalSalary >= 150000 {
		discount := (finalSalary * 0.1)
		fmt.Println("Su salario es igual a: ", finalSalary," Por lo cual se le descuenta: ", discount)
		finalSalary = finalSalary - discount
		fmt.Println("Por lo cual le queda en: ", finalSalary)
		fmt.Println("-------------------------------------------------------------------------------")
	} else {
		fmt.Println("Su salario es igual a: ", finalSalary," Al ser menor de 150.000 no paga impuesto")
		fmt.Println("-------------------------------------------------------------------------------")
	}
	worker.salaries = append( worker.salaries, finalSalary)
	return finalSalary, nil
}
func bonus(worker Worker) (float64, error)  {
	var (
		holder 		float64 = 0
		biggest 	float64 = 0
		finalBonus 	float64 = 0
		count 		float64 = 0
	)
	for _,values := range worker.salaries {
		count += 1
		if values > holder {
			holder = values
			biggest = holder
		}
	}
	if count < 3 {
		return 0, fmt.Errorf("No tiene meses suficientes para recibir aguinaldo, Meses trabajados: %v", count)
	}
	finalBonus = (biggest / 12) * count
	return finalBonus, nil
}