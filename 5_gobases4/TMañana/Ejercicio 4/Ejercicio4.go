package main

import(
	"fmt"
)

type errorDetail struct {
	Msg string
}

func (e *errorDetail) Error() string {
	return fmt.Sprintf("%v",  e.Msg)
}

func salaryCalc(valor int, horas int) (int, error) {
	salary := valor*horas
	if horas < 80{
		return 1, &errorDetail{
			Msg: "El trabajo no puede haber trabajado menos de 80 horas en el mes",
		}
	}
	if salary > 150000{
		salary = salary - (salary/100*10)
	}
	return salary, nil
	
}

func medioAguinaldo(mejor int, meses int) (int, error) {
	if horas > 150000{
		salary = salary - (salary/100*10)
	}
	aguinaldo := mejor / 12 * meses
	if horas < 80{
		return 1, &errorDetail{
			Msg: "El trabajo no puede haber trabajado menos de 80 horas en el mes",
		}
	
	return salary, nil
	
}
}

func main(){
	var salary int
	fmt.Scanln(&salary)
	total, err := salaryCalc(850, 96)

	if err != nil {
		fmt.Printf("Ha ocurrido un error: %v\n", err)
	}else{
		fmt.Printf("Su salario total es: %v", total)
	}


}