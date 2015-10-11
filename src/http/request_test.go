package http

import (
	"reflect"
	"testing"
)

func TestRequestParsing(t *testing.T) {
	request, _ := RequestFromString("GET /wiki/HTTP HTTP/1.0\r\nHost: ru.wikipedia.org\r\nTest1: Deleted\r\nTest1: test2\r\n\r\nthis is body")

	answerHeaders := Headers{
		"Test1": "test2",
		"Host":  "ru.wikipedia.org",
	}
	answerMethod := "GET"
	answerPath := "/wiki/HTTP"
	answerProto := "HTTP/1.0"

	if !reflect.DeepEqual(request.Headers, answerHeaders) {
		t.Error("Error parse headers")
	}

	if request.Method != answerMethod {
		t.Error("Error parse method")
	}

	if request.Path != answerPath {
		t.Error("Error parse path")
	}

	if request.Proto != answerProto {
		t.Error("Error parse proto")
	}
}
