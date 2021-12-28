package models

type Product struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Type  string  `json:"type" `
	Count int     `json:"count"`
	Price float64 `json:"price" `
	// Ciudad Ciudad  `json:"ciudad" `
}

// type Ciudad struct {
// 	Id     int    `json:"id"`
// 	Nombre string `json:"nombre" `
// }
