package smetana

import (
	"log"
	"strings"
	"testing"
)

func TestUsingCustomLogger(t *testing.T) {
	var target strings.Builder
	customLogger := log.New(&target, "", 0)
	SetLogger(customLogger)
	logger.Print("hello world")
	result := target.String()
	assertEqual(t, "hello world", result)
}
