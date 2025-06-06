package request

type Request struct {
	Path    string
	Method  string
	Headers map[string]string
	Body    string
	Meta    map[string]any
}

func New(path, method, body string, headers map[string]string) *Request {
	return &Request{
		Path:    path,
		Method:  method,
		Body:    body,
		Headers: headers,
		Meta:    make(map[string]any),
	}
}
