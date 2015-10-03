package config

import ()

var ()

type Config struct {
	Port     int    `json:port`
	Host     string `json:host`
	RootPath string `json:rootPath`
}
