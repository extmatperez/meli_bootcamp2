package internal

import (
	"fmt"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/18_storage2/TurnoManana/pkg/store"
)

type Transaction struct {
	ID       int    `json:"id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    string `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}

type Repository interface {
	GetAll() ([]Transaction, error)
	GetTransactionById(Id int) (Transaction, error)
	Store(id int, codigo, moneda, monto, emisor, receptor, fecha string) (Transaction, error)
	LastId() (int, error)
	Update(id int, codigo, moneda, monto, emisor, receptor, fecha string) (Transaction, error) //todos
	UpdateCodigoAndMonto(id int, codigo, monto string) (Transaction, error)
	Delete(id int) error
}

type repository struct {
	store store.Store
}

func NewRepository(db store.Store) Repository {

	return &repository{
		store: db,
	}
}

func (repo *repository) GetAll() ([]Transaction, error) {
	var transactions []Transaction
	err := repo.store.Read(&transactions)

	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (repo *repository) GetTransactionById(id int) (Transaction, error) {
	transactions, err := repo.GetAll()

	if err != nil {
		return Transaction{}, err
	}
	for _, t := range transactions {
		if t.ID == id {
			return t, nil
		}
	}
	return Transaction{}, fmt.Errorf("no existe la transaccion con id %d", id)
}

func (repo *repository) Store(id int, codigo, moneda, monto, emisor, receptor, fecha string) (Transaction, error) {

	tran := Transaction{id, codigo, moneda, monto, emisor, receptor, fecha}
	transactions, err := repo.GetAll()

	if err != nil {
		return Transaction{}, err
	}

	transactions = append(transactions, tran)

	err = repo.store.Write(transactions)

	if err != nil {
		return Transaction{}, err
	}

	return tran, nil
}

func (repo *repository) LastId() (int, error) {
	transactions, err := repo.GetAll()

	if err != nil {
		return 0, err
	}
	if len(transactions) == 0 {
		return 0, nil
	} else {
		return transactions[len(transactions)-1].ID, nil
	}

}

func (repo *repository) Update(id int, codigo, moneda, monto, emisor, receptor, fecha string) (Transaction, error) {

	transactions, err := repo.GetAll()

	if err != nil {
		return Transaction{}, err
	}

	tran := Transaction{id, codigo, moneda, monto, emisor, receptor, fecha}

	for i, t := range transactions {
		if t.ID == tran.ID {
			transactions[i] = tran
			err := repo.store.Write(transactions)
			if err != nil {
				return Transaction{}, err
			}
			return tran, nil
		}

	}
	fmt.Println("tran no encotrda", id)
	return Transaction{}, fmt.Errorf("Transaction %v no encontrada", id)

}

func (repo *repository) UpdateCodigoAndMonto(id int, codigo, monto string) (Transaction, error) {

	transactions, err := repo.GetAll()

	if err != nil {
		return Transaction{}, err
	}

	for i, t := range transactions {
		if t.ID == id {
			transactions[i].Codigo = codigo
			transactions[i].Monto = monto
			err := repo.store.Write(transactions)
			if err != nil {
				return Transaction{}, err
			}
			return transactions[i], nil
		}

	}
	return Transaction{}, fmt.Errorf("Transaction %v no encontrada", id)

}

func (repo *repository) Delete(id int) error {
	transactions, err := repo.GetAll()

	if err != nil {
		return err
	}

	for i, t := range transactions {
		if t.ID == id {
			transactions = RemoveIndex(transactions, i)
			err := repo.store.Write(transactions)
			if err != nil {
				return err
			}
			return nil
		}

	}
	return fmt.Errorf("no existe la transaccion con id: %v", id)
}

// func GetAllTransactionFromFolder() ([]Transaction,error){

// 	file, err:= ioutil.ReadFile(fileName)
// 	if(err != nil) {
// 		return nil,err
// 	}

// 	var transaction []Transaction

// 	err = json.Unmarshal([]byte(file), &transaction)

// 	if(err != nil) {
// 		return nil,err
// 	}

// 	return transaction,nil

// }

func RemoveIndex(s []Transaction, index int) []Transaction {
	return append(s[:index], s[index+1:]...)
}

// func UpdateJson(transactions [] Transaction) error{

// 	dataBytes, err := json.Marshal(transactions)
//     if err != nil {
// 		return err
//     }
// 	err = ioutil.WriteFile(fileName, dataBytes, 0644)
// 	if err != nil {
// 		return err
//     }

// 	return nil
// }
