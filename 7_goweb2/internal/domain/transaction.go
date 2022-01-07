/*
	Author: AG-Meli - Andrés Ghione
*/

package domain

type Transaction struct {
	ID       int     `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Remitter string  `json:"remitter"`
	Receptor string  `json:"receptor"`
	Date     string  `json:"date"`
}
