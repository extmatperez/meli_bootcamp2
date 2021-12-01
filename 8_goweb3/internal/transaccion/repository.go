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


	err = UpdateJson(transactions)
	


	if err != nil {
		return Transaction{},err
    }

	return tran,nil
}

func (repo *repository) LastId() (int, error) {
	transactions,err := GetAllTransactionFromFolder()

	if(err != nil){
		return 0,err
	}
	if len(transactions) == 0 {
		return 1,nil
	} else {
		return transactions[len(transactions)-1].ID ,nil
	}

}

func (repo *repository) Update(id int, codigo, moneda , monto, emisor, receptor,fecha string) (Transaction, error){

	transactions,err := GetAllTransactionFromFolder()

	if(err != nil){
		return Transaction{},err
	}

	tran := Transaction{id,codigo,moneda,monto,emisor,receptor,fecha}

	for i,t := range transactions {
		if t.ID == tran.ID {
			transactions[i] = tran
			UpdateJson(transactions)
			return tran,nil
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
			UpdateJson(transactions)
			return transactions[i],nil
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
			transactions = RemoveIndex(transactions,i)
			UpdateJson(transactions)
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


func UpdateJson(transactions [] Transaction) error{

	dataBytes, err := json.Marshal(transactions)
    if err != nil {
		return err
    }
	err = ioutil.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		return err
    }

	return nil
}