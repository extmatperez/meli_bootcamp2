package main

import ("fmt")

type myCustomError struct {
	status int
	msg string
}

// Approach 1
func (e *myCustomError) Error() string {
	return "Error desde el metodo"
} 

/*
// Approach 2
func (e *myCustomError) Error() string {
	return e.msg
	//Devuelve error desde la func
} 
*/

/* // Approach 3
func (e *myCustomError) Error() string {
	return fmt.Sprintf("Error desde %d - %v", e.status, e.msg)
} */

//Como esta funcion devuelve un error, buscara el metodo Error de la struct myCustomError y es lo que hacemos en los Approach 1 y 2.
func devolverError() error {
	var errorcito myCustomError
	//errorcito.msg = "Error desde la func"
	return &errorcito
}


/* Junto con el approach 3
func devolverError(asd int, asd2 string) error {
	var errorcito myCustomError
	errorcito.status = asd
	errorcito.msg = asd2
	return &errorcito
}
*/

func main() {

	err := devolverError()
	fmt.Println(err)

	/*
	err := devolverError(200, "asd")
	fmt.Println(err)
	*/

}