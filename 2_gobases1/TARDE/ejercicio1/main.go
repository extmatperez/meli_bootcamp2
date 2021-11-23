package main;

import "fmt";

func main(){
const palabra = "ejemplo";

fmt.Println("cantidad de letras: ",len(palabra))
for _, letra := range palabra {
	fmt.Println(string(letra))
}
}