package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configarations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables : ", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	servicePort := os.Getenv("SERVICE_NAME")
	if servicePort == "" {
		fmt.Println("Service Port is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("port must be number")
		os.Exit(1)
	}

	configarations = Config{
		Version:     version,
		ServiceName: servicePort,
		HttpPort:    int(port),
	}
}

func GetConfig() Config {
	loadConfig()
	return configarations
}
