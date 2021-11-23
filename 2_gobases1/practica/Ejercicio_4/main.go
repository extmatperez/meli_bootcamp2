package main

import "fmt"

func main(){

    // Correcto, la variable apellido esta asignada correctamente
    var apellido string = "Gomez"
    // Incorrecto, la variable edad se declara como int pero se inicializa como string, aqui sobraban las ""
    var edad int = 35
    // Incorrecto, sobran las comillas si se desea un boolean, además no es necesario el ;
    boolean := false
    // Incorrecto, la variable sueldo se declara string, pero se inicializa con numeros, se cambia a float64
    var sueldo float64 = 45857.90
    // Correcto.
    var nombre string = "Julián"

    fmt.Println(apellido," ", edad," ", boolean," ", sueldo," ", nombre)
}