package assert

import (
	"io"
	"os"
)

type location struct {
	Test     string
	FileName string
	Line     int
}

type errorLogger interface {
	Log(location *location, message string)
}

func newLogger() errorLogger {
	return &errorLoggerImpl{os.Stderr}
}

type errorLoggerImpl struct {
	writer io.Writer
}

func (logger *errorLoggerImpl) Log(location *location, message string) {
	// TODO: log using io.WriteString(logger.writer, ...)
}
