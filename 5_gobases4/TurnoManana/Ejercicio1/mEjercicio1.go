package main
import "fmt"


type MyError struct{

	status int
	msg string
}


func (e *MyError) Error() string {

	return fmt.Sprintf("%d - %v",e.status,e.msg)
}

func ProcesSalary(salary int) (error) {

	if salary < 150000 {
		return &MyError{
			status: 400,
			msg: "el salario ingresado no alcanza el mínimo imponible",
		}
	}
	return nil
}

func PagaImpuesto (salary int){
	err := ProcesSalary(salary) 

	if(err != nil) {
		fmt.Println(err)
	}else{

		fmt.Println("Debe pagar impuesto")
	}
}



func main() {

	PagaImpuesto(10)
	PagaImpuesto(210000)
 }


