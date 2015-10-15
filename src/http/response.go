package http

import (
	"bytes"
	"io/ioutil"
	"strconv"
	"time"
)

type Response struct {
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

func (r Response) Bytes() []byte {
	var result bytes.Buffer

	result.WriteString(r.Proto + " " + r.Status + stringSeparator)
	result.WriteString(r.Headers.ToPlainData() + stringSeparator)
	result.WriteString(r.Body)

	return result.Bytes()
}

func checkMethodErrors(response *Response, method string) bool {
	if !isCorrectMethod(method) {
		response.Status = statusInternalError
		return true
	}

	if !isSupportedMethod(method) {
		response.Status = statusMethodNotAllowed
		return true
	}

	return false
}

// TODO: rename it
func InitResponse(method string, path string) Response {
	response := Response{}
	response.Headers = Headers{}
	response.addDefaultHeaders()
	response.Proto = httpProto

	if checkMethodErrors(&response, method) {
		return response
	}

	if isDirectory(path) {
		path += defaultFile
	}

	data, err := ioutil.ReadFile(path)

	// TODO: Not only not found
	if err != nil {
		response.Status = statusNotFound
		return response
	}

	if method == "GET" {
		response.Body = string(data)
	}

	response.Status = statusOk
	response.Headers.Add("Content-type", getContentTypeByFilename(path))
	response.Headers.Add("Content-length", strconv.Itoa(len(data)))

	return response
}
