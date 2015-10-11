package http

import (
	"testing"
)

func TestHeaderDeserialize(t *testing.T) {
	headers := Headers{}

	answer := "Header1: 1\r\nHeader2: 2\r\n"
	headers.Add("Header1", "1")
	headers.Add("Header2", "3")
	headers.Add("Header2", "2")

	if headers.ToPlainData() != answer {
		t.Error("Error convert headers to plain data")
	}
}
