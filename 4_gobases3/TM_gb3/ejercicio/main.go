package main

import (
	"fmt"
	"os"
)

// Una empresa que se encarga de vender productos de limpieza necesita:
// Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
// Debe tener el id del producto, precio y la cantidad.
// Estos valores pueden ser hardcodeados o escritos en duro en una variable.

// Ejercicio 2 - Leer archivo
// La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

// Ejemplo:

// ID                            Precio  Cantidad
// 111223                      30012.00         1
// 444321                    1000000.00         4
// 434321                         50.50         1
//                           4030062.50

// type Product struct {
// 	Id           int     `json: "id"`
// 	Product_name string  `json: "product_name"`
// 	Price        float64 `json: "price"`
// 	Quantity     float64 `json: "quantity"`
// }

// func main() {

// 	product1 := Product{1, "escoba", 130.0, 3}
// 	product2 := Product{2, "lavandina", 80.5, 3}
// 	product3 := Product{3, "esponja", 20.5, 3}

// 	var list_products []Product

// 	list_products = append(list_products, product1, product2, product3)

// 	byte_list, err := json.Marshal(list_products)

// 	if err == nil {
// 		fmt.Println(byte_list)

// 		os.WriteFile("./products_stock.txt", byte_list, 0644)

// 	}

// 	data, err := os.ReadFile("./products_stock.txt")

// 	if err == nil {
// 		readFile := fmt.Println(json.Unmarshal(data, ))
// 	}

// }

type Product struct {
	Id       float64
	Price    float64
	Quantity float64
}

func main() {

	product1 := Product{111223, 30012.00, 1}
	product2 := Product{444321, 1000000.00, 4}
	product3 := Product{434321, 50.50, 1}

	var list_products []Product

	list_products = append(list_products, product1, product2, product3)

	fmt.Println(list_products)

	var productInfo string

	for _, product := range list_products {
		productInfo = fmt.Sprintf("%v;%v;%v\n", product.Id, product.Price, product.Quantity)
	}

	os.WriteFile("./products_stock.txt", []byte(productInfo), 0644)

}
