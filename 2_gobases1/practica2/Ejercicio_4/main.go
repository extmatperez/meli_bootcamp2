package  main

import (
	"fmt"
	"time"
)

func main(){

	var (
		number = 7
	)
	fmt.Println("el numero es: ", number, "el mes es: ", time.Month(number))
}