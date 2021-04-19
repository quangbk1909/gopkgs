package httpx

import "fmt"

type ErrorHTTPCall struct {
	Status int
	Body   []byte
}

func (e *ErrorHTTPCall) Error() string {
	return fmt.Sprintf("http error, status: %d, response: %s", e.Status, e.Body)
}
