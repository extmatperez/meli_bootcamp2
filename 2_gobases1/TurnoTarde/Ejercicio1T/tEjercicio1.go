package main
import "fmt"



func main() {
 var palabra string = "Hola soy una cadena";

 fmt.Printf("El tamaño de la palabra es de %v letras\n", len(palabra))

for i,letra := range palabra {
 
fmt.Printf("\nLetra en la posicion %d es %c \n", i,letra)

}
}
