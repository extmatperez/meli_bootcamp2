package main
import "fmt"
import "strings"

func main(){
	
	palabra := "prueba"
	fmt.Println("Cantidad de letras de la palabras es:",len(palabra))
	letraPorLetra := []string{}		
	letraPorLetra = strings.Split(palabra,"")	
	for i:=0;i<len(letraPorLetra);i++{
		fmt.Println(letraPorLetra[i])		
	}
}