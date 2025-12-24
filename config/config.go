package main

import (
	
)

type Logger struct {}

type Config struct {
	Host, Port string
}

func NewLogger() *Logger {
	return &Logger{}
}
