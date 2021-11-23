package main

import "fmt"

func main (){
    fmt.Println("hola Mundo!")
    // set variables 1
    var edad int
    edad = 27
    // print variables without format
    fmt.Println(edad)
    // print variables with format
    fmt.Printf("La edad es: %v, %d, %T", edad, edad, edad)
    // set variables 2
    horas := 20.5

    fmt.Printf("\n Las \" horas son; %v, %f, %T", horas, horas, horas)
}