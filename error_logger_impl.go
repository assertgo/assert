package assert

import (
	"fmt"
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

var theLogger errorLogger = &errorLoggerImpl{os.Stderr}

type errorLoggerImpl struct {
	writer io.Writer
}

func (logger *errorLoggerImpl) Log(location *location, message string) {
	fmt.Fprintf(logger.writer, "--- FAIL: %s\n\t%s:%d\n\t\t%s\n", location.Test, location.FileName, location.Line, message)
}
