package config

// Config model.
type Config struct {
	Port          int      `json:port`
	Host          string   `json:host`
	RootPath      string   `json:rootPath`
	LoggerEngines []string `json:loggerEngines`
}

// Constants for configuration logging.
const (
	FileConfig   = "FILE"
	StdoutConfig = "STDOUT"
)
