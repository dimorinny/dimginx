package http

import (
	"reflect"
	"testing"
)

func TestParsingHeaders(t *testing.T) {
	request, _ := RequestFromString("GET /wiki/HTTP HTTP/1.0\nHost: ru.wikipedia.org\nTest1: Deleted\nTest1: test2")
	answer := Headers{
		"Test1": "test2",
		"Host":  "ru.wikipedia.org",
	}

	if !reflect.DeepEqual(request.Headers, answer) {
		t.Error("Error parse headers")
	}
}
