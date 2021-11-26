package main

import(
	"fmt"
)

type errorDetail struct {
	Code int
	Msg string
}

func (e *errorDetail) Error() string {
	return fmt.Sprintf("%d - %v", e.Code, e.Msg)
}

func salaryCalc(salary int) error {
	if salary < 150000{
		return &errorDetail{
			Code: 100,
			Msg: "Su salario no alcanza el minimo imponible",
		}
	}
	return nil
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