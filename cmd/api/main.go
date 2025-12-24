package main

import (

)

func main() {
	// Create dependencies
	logger := NewLogger() // config pkg
	cfg := &Config{Host: "0.0.0.0", Port: "5000"} // config pkg

	Run(logger, cfg) // server pkg
}
