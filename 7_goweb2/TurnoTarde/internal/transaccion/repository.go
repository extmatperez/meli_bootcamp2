package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Transaction struct {
	ID       int   `json:"id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    string `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}

var lastID int
var fileName = "./transactions.json"

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, codigo, moneda , monto, emisor, receptor,fecha string) (Transaction, error)
	LastId() (int, error)
	Update(id int, codigo, moneda , monto, emisor, receptor,fecha string) (Transaction, error) //todos
	UpdateCodigoAndMonto(id int,codigo,monto string)(Transaction, error)
	Delete(id int)(error)
}

type repository struct {

}


func NewRepository() Repository{
	return &repository{}
}

func(repo *repository) GetAll() ([]Transaction, error){
	return GetAllTransactionFromFolder()

}

func (repo *repository) Store(id int, codigo, moneda , monto, emisor, receptor,fecha string) (	Transaction, error){
	
	tran := Transaction{id,codigo,moneda,monto,emisor,receptor,fecha}
	transactions,err := GetAllTransactionFromFolder()

	if(err != nil){
		return Transaction{},err
	}

	transactions = append(transactions, tran)

	dataBytes, err := json.Marshal(transactions)
    if err != nil {
		return Transaction{},err
    }


	err = ioutil.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		return Transaction{},err
    }

	return tran,nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}

func (repo *repository) Update(id int, codigo, moneda , monto, emisor, receptor,fecha string) (Transaction, error){

	transactions,err := GetAllTransactionFromFolder()

	if(err != nil){
		return Transaction{},err
	}

	tran := Transaction{id,codigo,moneda,monto,emisor,receptor,fecha}

	for i,t := range transactions {
		if t.ID == tran.ID {
			transactions[i] = t
			return t,nil

		}
		
	} 
	return Transaction{}, fmt.Errorf("Transaction %v no encontrada",id)

}


func (repo *repository) UpdateCodigoAndMonto(id int, codigo,monto string ) (Transaction, error){

	transactions,err := GetAllTransactionFromFolder()

	if(err != nil){
		return Transaction{},err
	}

	for i,t := range transactions {
		if t.ID == id {
			transactions[i].Codigo = codigo
			transactions[i].Monto = monto
			return t,nil
		}
		
	} 
	return Transaction{}, fmt.Errorf("Transaction %v no encontrada",id)

}


func (repo *repository) Delete(id int) error {
	transactions,err := GetAllTransactionFromFolder()

	if(err != nil){
		return err
	}

	for i,t := range transactions {
		if t.ID == id {
			RemoveIndex(transactions,i)
			return nil
		}
		
	} 
	return fmt.Errorf("no existe la transaccion con id: %v",id)
} 

func GetAllTransactionFromFolder() ([]Transaction,error){
	
	file, err:= ioutil.ReadFile(fileName)
	if(err != nil) {
		return nil,err
	}

	var transaction []Transaction
 
	err = json.Unmarshal([]byte(file), &transaction)
	
	if(err != nil) {
		return nil,err
	}

	return transaction,nil

}

func RemoveIndex(s []Transaction, index int) []Transaction {
    return append(s[:index], s[index+1:]...)
}