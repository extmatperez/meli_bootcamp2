package main
import "fmt"



func main() {
  var precio float64 = 10.0
  var descuento float64 = 0.2
precioFinal := precio * (1-descuento)
   
fmt.Printf("El precio del producto con el descuento de %v %% es de $ %v\n", 
descuento*100,precioFinal)




}
