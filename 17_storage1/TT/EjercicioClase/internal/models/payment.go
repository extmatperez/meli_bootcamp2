package models

type Payment struct {
	Id         int        `json:"id"`
	Codigo     string     `json:"codigo"`
	Moneda     string     `json:"moneda"`
	Monto      float64    `json:"monto"`
	Emisor     string     `json:"emisor"`
	Receptor   string     `json:"receptor"`
	Fecha      string     `json:"fecha"`
	BoxClosing BoxClosing `json:"box_closing"`
}

type BoxClosing struct {
	Id          int    `json:"id"`
	Responsable string `json:"responsable"`
	Fecha       string `json:"fecha"`
}
