package main

import "fmt"
import "errors"


func promedio(notas ...int) (int, error){
	var promedio int

	for _, nota := range notas{
		if nota == 0 {
			return 0, errors.New("Invalido, una de las notas es igual a 0")
		}
		promedio += nota
	}

	return promedio / len(notas), nil
}


func main() {
	promedio, err := promedio(5,9,10,8,4,8,10,10,0, 10)
	if err == nil{
		fmt.Println("Promedio del alumno:", promedio)
	}else{
		fmt.Println(err)
	}
}