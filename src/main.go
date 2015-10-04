package main

import (
	"config"
	"http"
	"io"
	"log"
	"logger"
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

func initLogger() {
	var writer io.Writer
	if len(configuration.LoggerEngines) != 0 {
		writer, closers = config.ParseLogEngines(configuration.LoggerEngines, logger.LogFileName)
		logger.Init(writer)
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
