package main
import "fmt"

func main() {
var temperatura float32 = 25.6;
var presionAtmosferica int= 30;
var presion float32=100; 
var txtTemperatura string = "La temperatura actual es: ";
var simboloGrado string = "°";
fmt.Println( txtTemperatura,temperatura,simboloGrado);
fmt.Println("La presion atomesferica es: ",presionAtmosferica, "Pa" );
fmt.Println("La presion actual:",presion,"Pa");
}
