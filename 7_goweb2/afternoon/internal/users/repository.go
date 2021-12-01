package internal

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	LastName string  `json:"last_name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Age      int     `json:"age" binding:"required"`
	Height   float64 `json:"height" binding:"required"`
	Active   bool    `json:"active" binding:"required"`
	Created  string  `json:"created" binding:"required"`
}

var users []User
var lastID int

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]User, error) {
	return users, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, name, lastName, email string, age int, height float64, active bool, created string) (User, error) {
	u := User{id, name, lastName, email, age, height, active, created}
	users = append(users, u)
	lastID = u.ID
	return u, nil
}
