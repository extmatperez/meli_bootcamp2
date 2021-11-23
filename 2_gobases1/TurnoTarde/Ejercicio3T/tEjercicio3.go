package main

import (
	"fmt"
)



func main() {
  const EDADMAX int=22
  const MINANTIGUEDADLABORAL int = 1
  var SUELDOSININTERES float32 = 100000.0
  var edad int = 23
  var esEmpleado bool = true
  var antiguedadLaboral int = 2 

  var sueldo float32 = 8800
  if(EDADMAX < edad  && esEmpleado && MINANTIGUEDADLABORAL < antiguedadLaboral) {
      if(SUELDOSININTERES > sueldo){
        fmt.Println("Entregar prestamo sin interes")
      }else{
        fmt.Println("Entregar prestamo con interes")
      }

    }else{
      fmt.Println("No entregar prestamo")
    }




}
