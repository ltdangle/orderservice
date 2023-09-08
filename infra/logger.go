package infra

import "fmt"

// Logger interface.
type Logger interface {
	Log(msg string)
}

// SimpleLogger Logger implementation.
type SimpleLogger struct{}

// NewSimpleLogger constructor.
func NewSimpleLogger() Logger{
	return SimpleLogger{}
}

// Log method.
func (l SimpleLogger) Log(msg string) {
	fmt.Println(msg)
}
