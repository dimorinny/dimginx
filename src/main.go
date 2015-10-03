package main

import (
	"config"
	"fmt"
)

var (
	configuration config.Config
)

func init() {
	parseConfig()
}

func parseConfig() {
	parser := config.JsonConfigParser{}
	configuration, err := parser.Parse()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(configuration)
	}
}

func main() {
}
