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

const (
	failOutput                = "--- FAIL: %s\n\t%s:%d\n\t\t%s\n"
	failOutputWithoutFailLine = "\t%s:%d\n\t\t%s\n"
)

func (logger *errorLoggerImpl) Log(location *location, message string) {
	args := []interface{}{location.Test, location.FileName, location.Line, message}
	if logger.alreadyWritten == false {
		fmt.Fprintf(logger.writer, failOutput, args...)
		logger.alreadyWritten = true
	} else {
		fmt.Fprintf(logger.writer, failOutputWithoutFailLine, args[1:]...)
	}
}
