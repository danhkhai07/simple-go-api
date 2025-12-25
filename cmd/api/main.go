package main

import (
	"first-api/config"
	"first-api/server"
)

func main() {
	logger := config.NewLogger()
	cfg := &config.Config{Host: "0.0.0.0", Port: "5000"} 

	server.Run(logger, cfg)
}
