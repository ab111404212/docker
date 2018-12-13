package log

import (
	"io"
	"sync"
)

type Logger struct {
	Writer io.Writer
	mu     *sync.Mutex
}

func (lg *Logger) SetOutput() error {
	return nil
}

func (lg *Logger) New() *Logger {
	return nil
}
