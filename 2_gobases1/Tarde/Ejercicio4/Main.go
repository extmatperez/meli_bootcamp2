package main

import "fmt"

func main()  {
	
	fmt.Println("Arranca el ejercicio 4")

	var mes uint
	mes = 10

	//FORMA PRACTICA
	fmt.Println("FORMA PRACTICA")
	var mesesPractico = [12]string {"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	if mes < 12 {
		fmt.Println("El mes numero", mes, "es",mesesPractico[mes-1])
	} else {
		fmt.Println("Numero no valido para el mes")
	}
	
	//FORMA NO PRACTICA
	fmt.Println("FORMA NO PRACTICA")
	switch mes {
	case 1:
		fmt.Println("ENERO")
	case 2:
		fmt.Println("FEBRERO")
	case 3:
		fmt.Println("MARZO")
	case 4:
		fmt.Println("ABRIL")
	case 5:
		fmt.Println("MAYO")
	case 6: 
		fmt.Println("JUNIO")
	case 7:
		fmt.Println("JULIO")
	case 8:
		fmt.Println("AGOSTO")
	case 9:
		fmt.Println("SEPTIEMBRE")
	case 10:
		fmt.Println("OCTUBRE")
	case 11: 
		fmt.Println("NOVIEMBRE")
	case 12:
		fmt.Println("DICIEMBRE")
	default:
		fmt.Println("MES NO VALIDO")
	}

}