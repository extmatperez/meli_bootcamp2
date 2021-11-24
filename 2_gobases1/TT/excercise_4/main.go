package main

import "fmt"

func main() {
	mes := [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}
	var numero int
	fmt.Println("ingresa el numero del mes")
	fmt.Scanln(&numero)

	for i := 0; i < 12; i++ {
		if i+1 == numero {
			fmt.Println(mes[i])
		}
	}
}

/* Se podria realizar el ejercicio con if anidados pero seria mas trabajo computacional y mas codigo
pero decidi mejor usar una estructura tipo array ya que son datos estaticos y solo una funcion donde busque el indice
correspondiente */
