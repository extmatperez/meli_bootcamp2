package web

type Response struct {
	StatusCode int `json:"status_code"`
	Content interface{} `json:"content"`
	MessageError string `json:"message_error"`
}

func NewResponse(statusCode int, content interface{}, messageError string) Response{
	if statusCode < 300 {
		return Response{statusCode, content, ""}
	}
	return Response{statusCode, nil, messageError}
}