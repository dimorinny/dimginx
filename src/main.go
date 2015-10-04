package main

import (
	"config"
	"fmt"
	"io"
	"log"
	"logger"
	"os"
)

var (
	configuration *config.Config
	logFile       *os.File
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
			logFile = createFileForAppend(logger.LogFileName)
			writers = append(writers, logFile)
		} else if v == config.StdoutConfig {
			writers = append(writers, os.Stdout)
		}
	}

	return io.MultiWriter(writers...)
}

func initLogger() {
	fmt.Println(configuration.LoggerEngines)
	if len(configuration.LoggerEngines) != 0 {
		logger.Init(log.New(createMultiWriter(configuration.LoggerEngines), "", 0))
	}

	logger.LogD("qweqweqweqe", "sdadasd")
	logger.LogD("12312qweqweqweqe", "sdadasd")
}

func init() {
	parseConfig()
	initLogger()
}

func main() {
	if logFile != nil {
		defer logFile.Close()
	}

}
