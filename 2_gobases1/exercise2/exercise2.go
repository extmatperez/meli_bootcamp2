// Exercise Clima
package main

import "fmt"

var temperature int
var humidity int
var press float64

func main() {
	temperature = 26
	humidity = 39
	press = 1012.9
	fmt.Printf("En Santiago de Chile se tiene que: \nTemperatura: %d ºC \nHumedad: %d %% \nPresion %f hPa ", temperature, humidity, press)
}

/* Segun la app de tiempo weather tanto la temperatura como la humedad la muestra como valores enteros, y son valores pequeños, los cuales
no deberian sobrepasar los 3 digitos, por lo que un int8 esta mas que suficiente:

int8        signed  8-bit integers (-128 to 127)
int16       signed 16-bit integers (-32768 to 32767)
int32       signed 32-bit integers (-2147483648 to 2147483647)
int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

*/
