package main

func main() {
	/*
		La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
		1- Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
		2- Luego imprimí cada una de las letras.
	*/
	word := "Hola"

	println("\nLa palabra tiene: ", len(word), " letras\n")
	/*
		for i := 0; i < len(word); i++ {
			println(word[i : i+1])
		}
	*/

	for _, letter := range word {
		println(string(letter))
	}
	/*
		for i := 0; i < len(word); i++ {
			println(string(word[i]))
		}
	*/
}
