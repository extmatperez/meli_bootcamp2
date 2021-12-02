package handler

type request struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}
type User struct {

}
