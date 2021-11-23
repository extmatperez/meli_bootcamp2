package main;

import "fmt";

func main(){
var empleado bool;
var edad, experiencia int

empleado = true;
edad = 24;
experiencia = 0;

switch {
case empleado == false: fmt.Println("usted tiene que tener empleo para solicitarlo");

case edad < 22 : fmt.Println("usted debe tener al menos 22 años");

case experiencia < 1: fmt.Println("debe tener al menos un año de experiencia")

default: fmt.Println("usted puede solicitar un prestamo")
}

}