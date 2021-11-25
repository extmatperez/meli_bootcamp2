/*Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo, para cumplir el
objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y si gana más de $150.000 se le
descontará además un 10%.
*/

package main

import (
	"fmt"
)

func tax(salary float64) {
	if salary > 50000 {
		salary = salary * 0.83
	} else if salary > 150000 {
		salary = salary * 0.73
	}
	fmt.Println("El sueldo del empleado es de ", salary)

}

func main() {
	var salary float64
	salary = 70000
	tax(salary)

}
