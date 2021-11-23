package main

import "fmt"

func main(){
    // App de Clima
    //declare variables
    var temperature int
    var humidity float64
    var pressure int
    // set data
    temperature = 28
    humidity = 26.5
    pressure = 1013
    // print data
    fmt.Println("Bienvenido al Tiempo.\nLa temperatura de hoy en Santiago es: ", temperature,"º",
        "\nLa humedad en el aire es de un: ", humidity,"%", "\nY la presión atmosferica es de: ", pressure, "hPa")
}