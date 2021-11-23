package main

import "fmt"

func main() {
	var temperatura, humedad int;
	var presion float32;
	//Le asigno int a la temperatura y a la humedad porque siempre son valores sin coma
	//En cambio, presion es un valor con coma por lo que le asigno float.
	temperatura = 23;
	humedad = 30;
	presion = 1017.2;

	fmt.Println("Ciudad de Moron");
	fmt.Println("La temperatura es:", temperatura);
	fmt.Println("La humedad es:", humedad, "%");
	fmt.Println("La presion es:", presion);

}