package ordenamiento

func Ordenar(numeros []int) []int{
	for i := 1; i < len(numeros); i++ {
		auxiliar := numeros[i]
		for j := i - 1; j >= 0 && numeros[j] > auxiliar; j-- {
		 numeros[j+1] = numeros[j]
		 numeros[j] = auxiliar
		}
	   }
	   return numeros
}