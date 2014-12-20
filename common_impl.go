package assert

import (
	"io"
	"os"
)

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
}

type location struct {
	Test     string
	FileName string
	Line     int
}

func logIfFalse(t TestingT, condition bool, format string, args ...interface{}) {
	if !condition {
		t.Errorf(format, args...)
	}
}
