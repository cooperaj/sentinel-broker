package logging

import (
	"fmt"
	"log"
	"os"
)

// Logger logging instance
var (
	Logger     *log.Logger
	NullLogger *log.Logger
)

// DevNull a dummy struct to capture unwanted logging output.
type DevNull struct{}

func (DevNull) Write(p []byte) (int, error) {
	return len(p), nil
}

// Create creates a configured instance of logger
func Create(name string) {
	Logger = log.New(os.Stderr, name, log.LstdFlags)
	NullLogger = log.New(&DevNull{}, name, log.LstdFlags)
}

// Logf Allows the output of formatted log strings
func Logf(s string, args ...interface{}) {
	if Logger == nil {
		return
	}
	Logger.Output(2, fmt.Sprintf(s, args...))
}
