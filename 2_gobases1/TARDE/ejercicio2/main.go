package main;

import "fmt";


func main(){
const precio float64 = 550.5;
const descuento float64 = 50;
var aux float64;

aux = descuento * precio / 100;

fmt.Println("el valor es: ", precio - aux)
}