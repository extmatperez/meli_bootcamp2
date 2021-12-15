package main

import (
	"fmt"
)

/* import (
	"fmt"
	"os"
)


type error interface {
	Error() string
}

type myError struct {
	status int
	msg    string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%d, %v", e.status, e.msg)
}

func myErrorTest(status int) (int, error) {
	if status >= 300 {
		return 400, &myError{
			status: status,
			msg:    "something go wrong",
		}
	}

	return 200, nil
}

func main() {

	status, err := myErrorTest(200)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Status %d funciona!!", status)
}
*/

type error interface {
	Error() string
}

type myError struct {
	msg    string
	salary float64
}

func (e *myError) Error() string {
	return fmt.Sprintf("%v", e.msg)
}

func checkPay(salary float64) error {
	if salary < 150000 {
		return &myError{
			salary: salary,
			msg:    "error, el salario no alcanza el minimo imponible",
		}

	}

	return nil
}

func main() {

	err := checkPay(160000)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("Debe pagar impuesto")

}
