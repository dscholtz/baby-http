package request

type Response struct {
	Status  int
	Headers map[string]string
	Body    string
}

func NewResponse(status int, body string) *Response {
	return &Response{
		Status:  status,
		Headers: make(map[string]string),
		Body:    body,
	}
}
