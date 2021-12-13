package web

type Response struct {
	Status int				`json:"status"`
	Content interface{} 	`json:"content,omitempty"`
	Error string 			`json:"error,omitempty"`
}


func NewResponse (status int, content interface{}, err string) Response{
	if(status < 300){
		return Response{Status: status,Content: content,Error: ""}
	}
		return Response{Status: status,Content: nil,Error: err}
}