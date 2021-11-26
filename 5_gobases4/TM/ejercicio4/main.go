/*
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
1)Desarrolla las funciones necesarias para permitir a la empresa calcular:
	a) Salario mensual de un trabajador según la cantidad de horas trabajadas.
		La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
		Dicha función deberá retornar más de un valor (salario calculado y error).
		En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el
		10% en concepto de impuesto.
		En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
		la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber
		trabajado menos de 80 hs mensuales”.
	b) Calcular el medio aguinaldo correspondiente al trabajador
		Fórmula de cálculo de aguinaldo:
		[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
		La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso
		de que se ingrese un número negativo.

2) Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”,
“fmt.Errorf()” y “errors.Unwrap()”.

No olvides realizar las validaciones de los retornos de error en tu función “main()”.
*/

package main

import (
	"errors"
	"fmt"
)

func calcularSalario(horasTrabajadas float64, valorHora float64) (float64, error) {

	if horasTrabajadas < 80 || horasTrabajadas < 0 {
		return 0, errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {

		salario := horasTrabajadas * valorHora

		if salario >= 150000 {
			salario = salario - salario*0.10
		}

		return salario, nil
	}
}

func calcularAguinaldo(mejorSalario float64, mesesTrabajados int) (float64, error) {

	if mesesTrabajados < 0 || mejorSalario < 0 {
		return 0, errors.New("error: se ingresó un número negativo")
	} else {
		return (mejorSalario / 12) * float64(mesesTrabajados), nil
	}
}

func main() {

	salario, err := calcularSalario(120, 100)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El salario es de $%2.f\n", salario)
	}

	aguinaldo, err := calcularAguinaldo(100000, 5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("El aguialdo es de $%2.f\n", aguinaldo)
	}
}
