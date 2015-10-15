package main

import (
	"config"
	"fmt"
	"http"
	"io"
	"log"
	"logger"
	"net"
	"runtime"
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

func initNumCpus() {
	if configuration.NumCpus > 0 {
		runtime.GOMAXPROCS(configuration.NumCpus)
	} else {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}

func init() {
	initConfig()
	initLogger()
	initNumCpus()
}

func main() {
	listenParams := fmt.Sprintf("%s:%d", configuration.Host, configuration.Port)
	logger.LogI("Server started on " + listenParams)

	for _, v := range closers {
		defer v.Close()
	}

	listener, err := net.Listen("tcp", listenParams)
	if err != nil {
		logger.LogE(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			logger.LogE(err)
		}

		go handleConnection(conn)
	}
}

func readRequestData(c net.Conn) (*http.Request, error) {
	buff := make([]byte, 1024)

	_, readErr := c.Read(buff)

	if readErr != nil {
		return nil, readErr
	}

	return http.RequestFromString(string(buff))
}

func handleConnection(c net.Conn) {
	logger.LogI("New connection from " + c.RemoteAddr().String())
	defer c.Close()

	request, err := readRequestData(c)

	if err != nil {
		logger.LogD("Error read data from socket")
		return
	}

	response := http.InitResponse(request.Method, configuration.RootPath+request.Path)
	c.Write(response.Bytes())
}
