package models

type Producto struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Count       int     `json:"count"`
	Price       float64 `json:"price"`
	WarehouseID int     `json:"warehouseId"`
}

type DTOProducto struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Count     int       `json:"count"`
	Price     float64   `json:"price"`
	Warehouse Warehouse `json:"warehouse"`
}

type Warehouse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Adress string `json:"adress"`
}
