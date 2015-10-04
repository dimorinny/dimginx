package http

import (
	"strings"
)

const (
	headerSeparator = ": "
)

var ()

type Headers map[string]string

func headersFromPlainData(data []string) Headers {
	headerData := []string{}
	headers := Headers{}

	for _, v := range data {
		headerData = strings.Split(v, headerSeparator)
		headers[headerData[0]] = headerData[1]
	}

	return headers
}
