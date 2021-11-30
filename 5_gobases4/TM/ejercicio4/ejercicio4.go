/*Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrolla las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar
el 10% en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede
haber trabajado menos de 80 hs mensuales”.
Calcular el medio aguinaldo correspondiente al trabajador
Fórmula de cálculo de aguinaldo:
[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso
de que se ingrese un número negativo.

Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando
“errors.New()”, “fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de
los retornos de error en tu función “main()”.
*/

package main

import (
	"errors"
)

func salario(horas int, valor float64) (float64, error) {
	sueldo := 0.0

	if horas < 80 || horas < 0 {

		err := errors.New("el trabajador no puede haber trabajado menos de 80 hs mensuales")

		return 0, err
	} else {
		sueldo = valor * float64(horas)
	}

	if sueldo > 150000 {
		sueldo = sueldo - sueldo*0.10
	}

	return sueldo, nil
}

func aguinaldo(maxSalario float64, mesesTrabajado int) (float64, error) {

	if maxSalario < 0 {
		err := errors.New("El salario no puede ser negativo")

		return 0, err
	}

	if mesesTrabajado < 0 {

		err := errors.New("Los meses trabajados no pueden ser negativos")

		return 0, err
	}

	return (maxSalario / 12) * float64(mesesTrabajado), nil
}

func main() {

}
