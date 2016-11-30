package creator

type Response struct {
}

// DefaultResponse provides a default response by best effort.
func DefaultResponse() *Response {
	return &Response{}
}
