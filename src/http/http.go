package http

import (
	"strings"
)

const (
	stringSeparator  = "\r\n"
	requestSeparator = "\r\n\r\n"
	serverName       = "DimGINX"
	defaultFile      = "/index.html"
	notFoundFile     = "/404.html"
	httpProto        = "HTTP/1.1"
)

var (
	supportedMethods = []string{"GET", "HEAD"}
	httpMethods      = []string{"OPTIONS", "GET", "HEAD", "POST", "PUT", "PATCH",
		"DELETE", "TRACE", "CONNECT"}
)

var exts = map[string]string{
	"txt":  "application/text",
	"html": "text/html",
	"json": "application/json",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"js":   "text/javascript",
	"css":  "text/css",
	"gif":  "image/gif",
	"swf":  "application/x-shockwave-flash",
}

func getContentTypeByFilename(fileName string) string {
	parts := strings.Split(fileName, ".")
	return exts[parts[len(parts)-1]]
}

func isDirectory(fileName string) bool {
	return fileName[len(fileName)-1:] == "/"
}

func isCorrectMethod(method string) bool {
	for _, v := range httpMethods {
		if method == v {
			return true
		}
	}

	return false
}

func isSupportedMethod(method string) bool {
	for _, v := range supportedMethods {
		if method == v {
			return true
		}
	}

	return false
}

func isSecurePath(path string) bool {
	return !strings.Contains(path, "../")
}
