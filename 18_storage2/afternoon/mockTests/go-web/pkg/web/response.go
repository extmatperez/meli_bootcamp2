package web

type Response struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
}

func NewResponse(statusCode int, data interface{}, err string) *Response {
	if statusCode < 300 {
		return &Response{
			StatusCode: statusCode,
			Data:       data,
		}
	} else {
		return &Response{
			StatusCode: statusCode,
			Error:      err,
		}
	}
}
