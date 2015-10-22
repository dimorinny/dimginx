package http

import (
	"bytes"
	"io/ioutil"
	"os"
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
		response.Status = StatusInternalError
		return true
	}

	if !isSupportedMethod(method) {
		response.Status = StatusMethodNotAllowed
		return true
	}

	return false
}

func checkPathErrors(response *Response, path string) bool {
	if !isSecurePath(path) {
		response.Status = StatusForbidden
		return true
	}

	return false
}

func responseStatusByError(err error, isDirectory bool) string {
	var responseStatus string

	if os.IsNotExist(err) {
		if isDirectory {
			responseStatus = StatusForbidden
		} else {
			responseStatus = StatusNotFound
		}
	} else if os.IsPermission(err) {
		responseStatus = StatusForbidden
	}

	return responseStatus
}

func InitResponse(method string, path string) Response {
	response := Response{}
	response.Headers = Headers{}
	response.addDefaultHeaders()
	response.Proto = httpProto

	if checkMethodErrors(&response, method) || checkPathErrors(&response, path) {
		return response
	}

	isDirectoryFlag := isDirectory(path)

	if isDirectoryFlag {
		path += defaultFile
	}

	data, err := ioutil.ReadFile(path)

	// TODO: Not only not found
	if err != nil {
		response.Status = responseStatusByError(err, isDirectoryFlag)
		return response
	}

	if method == "GET" {
		response.Body = string(data)
	}

	response.Status = StatusOk
	response.Headers.Add("Content-Type", getContentTypeByFilename(path))
	response.Headers.Add("Content-Length", strconv.Itoa(len(data)))

	return response
}

func InitResponseForError(status string) Response {
	response := Response{}
	response.Headers = Headers{}
	response.addDefaultHeaders()
	response.Proto = httpProto
	response.Status = status
	return response
}
