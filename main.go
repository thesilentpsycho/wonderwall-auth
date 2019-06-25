package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/libertywireless/wonderwall-auth/config"
)

const _DEFAULT_CONFIG_PATH = "config.yml"

var configFilePath string
var serverPort uint

func main() {
	readFlags()
	initConfig()

	fmt.Println("Server listening on port :", serverPort)
	http.ListenAndServe(":"+strconv.Itoa(int(serverPort)), GetHandler())
}

func readFlags() {

	flag.StringVar(&configFilePath, "config", "config.yml", "Path to the config file")
	flag.UintVar(&serverPort, "server_port", 8080, "Port to run the server on")
	flag.Parse()
}

func initConfig() {

	if configFilePath == "" {
		configFilePath = _DEFAULT_CONFIG_PATH
	}
	config.LoadConfig(configFilePath)
}
