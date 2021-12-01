package internal

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

var users []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, first_name string, last_name string, age int) (User, error)
	// Store2(newUser User) ([]User, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]User, error) {
	return users, nil
}

func (repo *repository) Store(id int, first_name string, last_name string, age int) (User, error) {
	us := User{id, first_name, last_name, age}
	users = append(users, us)
	lastID = us.ID
	return us, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}
