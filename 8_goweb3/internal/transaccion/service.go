package internal

type Service interface{
	GetAll() ([]Transaction, error)
	GetTransactionById(Id int) (Transaction, error)
	Store(codigo, moneda , monto, emisor, receptor,fecha string) (Transaction, error)
	Update(id int, codigo, moneda , monto, emisor, receptor,fecha string) (Transaction, error) //todos
	UpdateCodigoAndMonto(id int,codigo,monto string)(Transaction, error)
	Delete(id int)(error)
}

type service struct{
	repository Repository
}

func NewService(repo Repository) Service{
	return &service{repo}
}


func (ser *service) GetAll() ([]Transaction, error) {

	transactions,err := ser.repository.GetAll()

	if(err != nil){
		return nil,err
	}

	return transactions,nil
}



func (ser *service) GetTransactionById(id int) (Transaction, error) {

	transaction,err := ser.repository.GetTransactionById(id)

	if(err != nil){
		return Transaction{},err
	}

	return transaction,nil
}



func (ser *service) Store(codigo, moneda , monto, emisor, receptor,fecha string) (Transaction, error) {
	
	lastID,err := ser.repository.LastId()

	if(err != nil){
		return Transaction{},err
	}
	
	lastID++
	transcation, err := ser.repository.Store(lastID,codigo,moneda,monto,emisor,receptor,fecha)
	
	if(err != nil){
		return Transaction{},err
	}
	return transcation,nil
}


func (ser *service) Update(id int,codigo, moneda, monto, emisor, receptor,fecha string) (Transaction, error) {
	
	transcation, err := ser.repository.Update(id,codigo,moneda,monto,emisor,receptor,fecha)
	
	if(err != nil){
		return Transaction{},err
	}
	return transcation,nil
}

func (ser *service) UpdateCodigoAndMonto(id int,codigo,monto string)(Transaction, error) {
	
	transcation, err := ser.repository.UpdateCodigoAndMonto(id,codigo,monto)
	
	if(err != nil){
		return Transaction{},err
	}
	return transcation,err
}
func (ser *service) Delete(id int)(error) {
	
	 err := ser.repository.Delete(id)
	
	if(err != nil){
		return err
	}
	return nil
}