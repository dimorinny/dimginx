package http

import (
	"bytes"
	"strings"
)

const (
	headerSeparator = ": "
)

type Headers map[string]string

func (h Headers) Add(key string, value string) {
	h[key] = value
}

func (h Headers) Get(key string) string {
	return h[key]
}

func (h Headers) ToPlainData() string {
	var result bytes.Buffer

	for k, v := range h {
		result.WriteString(k + headerSeparator + v + stringSeparator)
	}

	return result.String()
}

func headersFromPlainData(data []string) Headers {
	headerData := []string{}
	headers := Headers{}

	for _, v := range data {
		headerData = strings.Split(v, headerSeparator)

		if len(headerData) >= 2 {
			headers.Add(headerData[0], headerData[1])
		}
	}

	return headers
}
