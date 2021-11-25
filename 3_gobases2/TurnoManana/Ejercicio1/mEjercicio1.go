package main
import "fmt"

func main() {
	salario := getImpuestoSalary(25000)
	salario1 := getImpuestoSalary(51000)
	salario2 := getImpuestoSalary(151000)

	fmt.Println(salario,salario1,salario2)
 }


func getImpuestoSalary(salary float64) float64{
	const SALARYDISCCOUNT17 float64 = 50000.00
	const SALARYDISCCOUNT10 float64 = 150000.00
	if(salary > SALARYDISCCOUNT17 && salary < SALARYDISCCOUNT10){
		return salary * (1 - 0.17)
	}else if(salary > SALARYDISCCOUNT10){
		return salary * (1- 0.1)
	}
	return salary


}
