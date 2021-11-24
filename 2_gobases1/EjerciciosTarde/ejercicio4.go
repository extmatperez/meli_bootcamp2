package main

import "fmt"

//Punto 1
func main() {

	var meses = map[int]string { 1 : "Enero", 2 : "Febrero", 3 : "Marzo", 4 : "Abril", 5: "Mayo", 6 : "Junio", 7 : "Julio" , 8 : "Agosto", 9 : "Septiembre", 10 : "Octubre" , 11 : "Noviembre", 12 : "Diciembre"  }

	var mes int
	mes = 11

	if (mes > 0 && mes < 13){
		fmt.Println(meses[mes])
	}else{
		fmt.Println("Numero de mes no valido")
	}
}
//Punto 2
// Se podria hacer con un switch case, tambien con If anidados y en mapa
// Considero que es mejor en mapa ya que cada numero de mes tiene un valor en texto asociado a el y en caso de necesitar hacer otra validacion puedo reusar el mapa declarado