package main

/*
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrolla las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
Calcular el medio aguinaldo correspondiente al trabajador
Fórmula de cálculo de aguinaldo:
[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número negativo.

Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.
*/

import (
	"errors"
	"fmt"
	"os"
)

func halfSalaryTest(best_salary float64, months int) (float64, error) {
	var half float64
	if months <= 0 {
		return 0, errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	}
	half = (best_salary * float64(months)) / 12
	return half, nil
}

func myCustomErrorTest(hours float64, pricePerHour float64) (float64, error) {
	var salary float64 = hours * pricePerHour
	if hours < 0 || hours < 80 {
		return 0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	if hours >= 80 && salary >= 150000 {
		salary = salary * 0.9
	}
	return salary, nil
}

func main() {
	salary, err1 := myCustomErrorTest(80.00, 2500.00)
	half, err2 := halfSalaryTest(180000, 6)
	if err1 != nil || err2 != nil {
		if err1 != nil {
			fmt.Println(err1.Error())
		}
		if err2 != nil {
			fmt.Println(err2.Error())
		}
		os.Exit(1)
	}
	fmt.Println("Sueldo", salary)
	fmt.Println("Aguinaldo", half)
}
