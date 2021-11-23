package main 

import "fmt"

func main () {
	var temperatura int =24
	var humedad float32 =0.67
	var presion int =560
	
	fmt.Println(temperatura)
	fmt.Println(humedad)
	fmt.Println(presion)
	fmt.Printf("Temperatura: %v %T Humedad:%v %T  Presion:%v %T \n",temperatura,temperatura,humedad,humedad,presion,presion)

}