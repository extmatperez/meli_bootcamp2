package internal

import (
	"encoding/json"
	"io/ioutil"
)

type Transaccion struct {
	ID       int   `json:"id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    string `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}

var lastID int
var transacciones []Transaccion
var fileName = "./transactions.json"

type Repository interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigo, moneda , monto, emisor, receptor,fecha string) (Transaccion, error)
	LastId() (int, error)
}

type repository struct {

}


func NewRepository() Repository{
	return &repository{}
}

func(repo *repository) GetAll() ([]Transaccion, error){

	file, err1:= ioutil.ReadFile(fileName)

	if(err1 != nil) {
		return nil,err1
	}

	var transaction []Transaccion

	err := json.Unmarshal([]byte(file), &transaction)
	
	if(err != nil) {
		return nil,err
	}

	return transaction,nil

}

func (repo *repository) Store(id int, codigo, moneda , monto, emisor, receptor,fecha string) (Transaccion, error){
	
	tran := Transaccion{id,codigo,moneda,monto,emisor,receptor,fecha}
	
	transacciones = append(transacciones, tran)

	dataBytes, err3 := json.Marshal(transacciones)
    if err3 != nil {
		return Transaccion{},err3
    }


	err4 := ioutil.WriteFile(fileName, dataBytes, 0644)
	if err4 != nil {
		return Transaccion{},err4
    }

	return tran,nil
}


func (repo *repository) LastId() (int, error) {
	return lastID, nil
}

