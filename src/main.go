package main

import (
	"config"
	"flag"
	"fmt"
	"http"
	"io"
	"location"
	"logger"
	"net"
	"runtime"
)

var (
	ncpu            int
	configuration   *config.Config
	locationManager location.LocationManager
	closers         []io.Closer
)

// *** Args *** //

func initCommandLineArgs() {
	flag.IntVar(&ncpu, "c", runtime.NumCPU(), "The number of cpu")
	flag.Parse()
}

// *** Configuration *** //

func initConfig() {
	parser := config.JsonConfigParser{}
	conf, err := parser.Parse()

	configuration = conf

	if err != nil {
		logger.LogE(err)
	}
}

// *** Locations *** //

func initLocations() {
	var bufferLocations []*location.Location

	for _, v := range configuration.Locations {
		bufferLocation, err := location.InitLocation(v.Rule, v.Root)
		if err != nil {
			logger.LogE(err)
		}
		bufferLocations = append(bufferLocations, bufferLocation)
	}

	locationManager = location.InitLocationManager(bufferLocations)
}

// *** Logging *** //

func initLogger() {
	var writer io.Writer
	if len(configuration.LoggerEngines) != 0 {
		writer, closers = config.ParseLogEngines(configuration.LoggerEngines,
			logger.LogFileName)
		logger.Init(writer)
	}
}

func initNumCpus() {
	runtime.GOMAXPROCS(ncpu)
}

func init() {
	initCommandLineArgs()
	initConfig()
	initLogger()
	initLocations()
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

func generateResponse(method string, path string) http.Response {
	prefix, err := locationManager.Match(path)

	if err != nil {
		// TODO: maybe another error
		return http.InitResponseForError(http.StatusNotFound)
	}

	return http.InitResponse(method, prefix.Root+path)
}

func handleConnection(c net.Conn) {
	logger.LogI("New connection from " + c.RemoteAddr().String())
	defer c.Close()

	request, err := readRequestData(c)

	if err != nil {
		return
	}

	response := generateResponse(request.Method, request.Path)
	c.Write(response.Bytes())
}
