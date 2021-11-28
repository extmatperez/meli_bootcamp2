/*
? Ejercicio 4 -  Impuestos de salario #4
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrolla las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).

En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en
concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado
menos de 80 hs mensuales”.
Calcular el medio aguinaldo correspondiente al trabajador
Fórmula de cálculo de aguinaldo:
[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese
un número negativo.

Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”,
“fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.

*/

package main

import "fmt"

func salarioMensual(horas, valor float64) (float64, error) {
	salario := horas * valor

	if horas <= 80 || horas < 0 {
		return 0.0, fmt.Errorf("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	if salario >= 150000 {
		salario = salario - salario*0.1
		return salario, nil
	}
	return salario, nil
}

func aguinaldo(mejorSueldo float64, mesesTrabajados int) (float64, error) {
	if mejorSueldo <= 0 || mesesTrabajados <= 0 {
		return 0, fmt.Errorf("error: Los datos ingresados no pueden ser iguales o menores a 0 ")
	}
	return 1, nil
}

func main() {
	horas := 80.0
	valor := 1000.0
	result, err := salarioMensual(horas, valor)

	fmt.Println(result, err)

	mejorSueldo := 200000.0
	mesesTrabajados := 12

	agui, err1 := aguinaldo(mejorSueldo, mesesTrabajados)

	fmt.Println(agui, err1)

}
