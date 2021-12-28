package models

type Wharehouse struct {
	ID   int    `json:"id"`
	Name string `json:"nombre"`
}

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Price float64 `json:"price"`
	Size  int     `json:"size"`

	// Location Wharehouse `json:"location"`
}

// type DTOAvgAge struct {
// 	Name  string  `json:"name"`
// 	Average float64 `json:"average"`
// }
