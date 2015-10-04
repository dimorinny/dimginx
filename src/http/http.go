package http

import (
	"errors"
	"net/url"
	"strings"
)

var ()

const (
	stringSeparator = "\n"
)

type Request struct {
	Url     *url.URL
	Headers Headers
	Method  string
}

func (r *Request) parseHeaders(headers []string) {
	r.Headers = headersFromPlainData(headers)
}

func (r *Request) parseStartingLine(headers string) {}

func RequestFromString(req string) (*Request, error) {
	request := new(Request)
	parsedReq := strings.Split(req, stringSeparator)

	if len(parsedReq) == 0 {
		return nil, errors.New("Empty request")
	}

	// TODO: parse request body too
	request.parseStartingLine(parsedReq[0])
	request.parseHeaders(parsedReq[1:])

	return request, nil
}
