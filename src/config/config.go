package config

import (
	"io"
	"os"
)

// Config model.
type Config struct {
	Port          int      `json:port`
	Host          string   `json:host`
	RootPath      string   `json:rootPath`
	LoggerEngines []string `json:loggerEngines`
	NumCpus       int      `json:numCpus`
}

// Constants for configuration logging.
const (
	FileEngine   = "FILE"
	StdoutEngine = "STDOUT"
)

// Tools

func ParseLogEngines(writerStrings []string, logFileName string) (io.Writer, []io.Closer) {
	writers := []io.Writer{}
	closers := []io.Closer{}

	for _, v := range writerStrings {

		if v == FileEngine {
			logFile := createFileForAppend(logFileName)
			closers = append(closers, logFile)
			writers = append(writers, logFile)
		} else if v == StdoutEngine {
			writers = append(writers, os.Stdout)
		}
	}

	return io.MultiWriter(writers...), closers
}

func createFileForAppend(fileName string) *os.File {
	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return file
}
