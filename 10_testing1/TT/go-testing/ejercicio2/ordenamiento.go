package ordenamiento

func Ordenar(lista []int) []int {
	for i := 0; i < len(lista); i++ {
		for j := i + 1; j < len(lista); j++ {
			if lista[i] > lista[j] {
				lista[i], lista[j] = lista[j], lista[i]
			}
		}
	}
	return lista
}
