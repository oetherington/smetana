package smetana

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", 0)

func SetLogger(newLogger *log.Logger) {
	logger = newLogger
}
