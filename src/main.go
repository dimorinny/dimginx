package main

import (
	"config"
	"http"
	"io"
	"log"
	"logger"
	"os"
)

var (
	configuration *config.Config
	closers       []io.Closer
)

// *** Configuration *** //

func parseConfig() {
	parser := config.JsonConfigParser{}
	conf, err := parser.Parse()

	configuration = conf

	if err != nil {
		log.Fatal(err)
	}
}

// *** Logging *** //

func createFileForAppend(fileName string) *os.File {
	file, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return file
}

func createMultiWriter(writerStrings []string) io.Writer {
	writers := []io.Writer{}

	for _, v := range writerStrings {

		if v == config.FileConfig {
			logFile := createFileForAppend(logger.LogFileName)
			closers = append(closers, logFile)
			writers = append(writers, logFile)
		} else if v == config.StdoutConfig {
			writers = append(writers, os.Stdout)
		}
	}

	return io.MultiWriter(writers...)
}

func initLogger() {
	if len(configuration.LoggerEngines) != 0 {
		logger.Init(log.New(createMultiWriter(configuration.LoggerEngines), "", 0))
	}
}

func init() {
	parseConfig()
	initLogger()
}

func main() {
	for _, v := range closers {
		defer v.Close()
	}

	// Tests.
	request, _ := http.RequestFromString("GET /wiki/HTTP HTTP/1.0\nHost: ru.wikipedia.org\nLoh: Dimon")

	for k, v := range request.Headers {
		logger.LogD(k)
		logger.LogD(v)
	}
}
