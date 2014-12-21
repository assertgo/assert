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

var theLogger errorLogger = &errorLoggerImpl{writer: os.Stdout}

type errorLoggerImpl struct {
	writer         io.Writer
	alreadyWritten bool
}

func (logger *errorLoggerImpl) Log(location *location, message string) {
	if logger.alreadyWritten == false {
		fmt.Fprintf(logger.writer, "--- FAIL: %s\n\t%s:%d\n\t\t%s\n", location.Test, location.FileName, location.Line, message)
		logger.alreadyWritten = true
	} else {
		fmt.Fprintf(logger.writer, "\t%s:%d\n\t\t%s\n", location.FileName, location.Line, message)
	}
}
