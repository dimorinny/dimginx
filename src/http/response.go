package http

import (
	"time"
)

type Response struct {
	Method  string
	Status  string
	Proto   string
	Body    string
	Headers Headers
}

// Headers must be inited
func (r *Response) addDefaultHeaders() {
	r.Headers.Add("Server", serverName)
	r.Headers.Add("Date", time.Now().Format(dateTimeFormat))
	r.Headers.Add("Connection", "close")
}

func InitResponse(method string, path string) (*Response, error) {
	response := Response{}
	response.Headers = Headers{}

	// TODO: main generate response here

	return &response, nil
}
