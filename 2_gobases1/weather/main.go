package main

import "fmt"

func main(){
	var temperature float64 = 23.5
	var humidity float64 = 40
	var presion float64 = 50

	fmt.Printf("The temperature is: %v %T \n", temperature, temperature)
	fmt.Printf("The humidity is: %v %% %T \n", humidity, humidity)
	fmt.Printf("The presion is: %v %% %T \n", presion, presion)
}