package main;

import "fmt";

func main(){
	var temperatura, presion int;
	var humedad float64;
	temperatura = 24;
	humedad = 45.6;
	presion = 1050;

	fmt.Println("temperatura: ", temperatura, "ºC");
	fmt.Println("humedad: ", humedad, "%");
	fmt.Println("presion: ", presion)

}

//las declararia como float a las variables que puedan llegar a ser numeros con coma e int a las que sean numeros enteros.