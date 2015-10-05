package http

import (
	"errors"
	// "fmt"
	// "net/url"
	"strings"
)

var ()

const (
	stringSeparator  = "\r\n"
	requestSeparator = "\r\n\r\n"
)

type Request struct {
	Method  string
	Path    string
	Proto   string
	Body    string
	Headers Headers
}

func (r *Request) parseHeaders(headers []string) {
	r.Headers = headersFromPlainData(headers)
}

func (r *Request) parseStartingLine(startingLine string) {
	parts := strings.Split(startingLine, " ")
	uri := strings.Split(parts[1], "?")
	path, _ := url.QueryUnescape(uri[0])

	// TODO: parse query params
	r.Proto = parts[2]
	r.Path = uri[0]
	r.Method = parts[0]
}

func RequestFromString(req string) (*Request, error) {
	var headerReq string

	request := new(Request)
	headerReq, request.Body = splitRequest(req)
	parsedHeader := strings.Split(headerReq, stringSeparator)

	if len(parsedHeader) == 0 {
		return nil, errors.New("Empty request")
	}

	request.parseStartingLine(parsedHeader[0])
	request.parseHeaders(parsedHeader[1:])

	return request, nil
}

func splitRequest(request string) (string, string) {
	result := strings.Split(request, requestSeparator)

	if len(result) == 1 {
		return result[0], ""
	}

	return result[0], result[1]
}
