package main

import (
	"fmt"
)

type myError struct {
	message string
	code    int
}

func (e *myError) Error() string {
	return e.message
}

func errorControl(salary int) (int, error) {

	if salary < 150000 {
		return 0, &myError{
			message: "error: the salary entered does not reach the minimum available",
			code:    400,
		}
	}
	return salary, nil
}

func main() {
	var salary int = 15000
	response, err := errorControl(salary)

	if err != nil {
		error := fmt.Errorf("Real message: %s", err)
		fmt.Println("Get error from function:", error)
	} else {
		fmt.Println("The salary available is: ", response)
	}
}
