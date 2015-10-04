package main

import (
	"config"
	"fmt"
	"io"
	"log"
	"logger"
)

var (
	configuration *config.Config
	closers       []io.Closer
)

// *** Configuration *** //

func initConfig() {
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
	initConfig()
	initLogger()
}

func main() {
	logger.LogI(fmt.Sprintf("Server started on %s:%d", configuration.Host, configuration.Port))

	for _, v := range closers {
		defer v.Close()
	}

	// Server.
}
