package main

import "fmt"

func main() {
	x, y := 10, 20

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %v\n", x/y)
	fmt.Printf("x mod (resto) y = %d\n", x%y)

	// Array y Slices. Todos los elementos son del mismo tipo de dato.
	// Si tiene un tamaño determinado es un array.
	var a [2]string
	a[0] = "Rodrigo"
	a[1] = "Vega Gimenez"

	fmt.Println(a)

	// Si tiene tamaño sin definir es variable, y es un slice.
	var b = []bool{true, false, true, false, false}
	fmt.Println(b[0])
	fmt.Printf("\n%v", b)
	// Tambiem podemos definir rangos. Se puede hacer sobre arrays y slices.
	fmt.Println(b[1:4], b[1:], b[:3])

	fmt.Printf("\nLongitud de b: %d, cap: %d", len(b), cap(b))

	// Cuando se crean slices con make, y nos pasamos de la capacidad, siempre se duplican de ahi en mas, por ejemplo, llena las 8 y no llega a soportarlo, el nuevo tamaño del slice sera 16.
	c := make([]int, 6)
	fmt.Println("\n", c)
	extra_slice := make([]int, 3)
	c = append(c, extra_slice[0])
	// Aqui veremos que como ya tenemos en C los 6 elementos 0, si agregamos otro elemento el slice se dobla en tamaño y su capacidad es de 12, pero solo posee 7 elementos.
	fmt.Printf("\nLongitud de c: %d, cap: %d", len(c), cap(c))

	// Maps. Nos permiten crear variables de tipo clave-valor, tipo los diccionarios, definiendo un tipo de dato para las claves y otro para los valores.
	var my_map = map[string]int{}
	// Otra forma de llamarlo es my_map2 := make(map[string]string)
	fmt.Printf("\nMy map %v %T", my_map, my_map)
	fmt.Printf("\nMy map en posición [Rodrigo] %v", my_map["Rodrigo"])
	my_map["Rodrigo"] = 25
	fmt.Printf("\nMy map en posición [Rodrigo] %v, ahora lo deberia mostrar.", my_map["Rodrigo"])
	my_map["Luciana"] = 24
	my_map["Gerardo"] = 23
	my_map["Matias"] = 25
	my_map["Abel"] = 26
	my_map["Rodrigo"] = 38
	// Vemos que se sobreescribe el valor en el mapa.
	fmt.Println("\n", my_map["Rodrigo"])

	// Un ciclo for nos permite recorrer los elementos de un map.
	for key, element := range my_map {
		fmt.Printf("\nKey: %v, value: %v", key, element)
	}

	// GO solo tiene la estructura repetitiva FOR, nada mas.
	// Se usa como el for clasico, for range, un bucle infinito o un bucle wh¡le.
	// El for clasico es el siguiente
	for i := 0; i < len(b); i++ {
		i += i
		fmt.Printf("\nValores de b: %v", b[i])
	}

	// El for range es para iterar sobre estructuras de datos, por ejemplo, tenemos el array b creado anteriormente.
	for _, element := range b {
		fmt.Printf("\n%v", element)
	}

	// IF...ELSE.
	// SWITCH...

}
