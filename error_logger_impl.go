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
	writer       io.Writer
	prevTestName string
	prevTestLine int
}

const (
	failOutput                  = "--- FAIL: %s\n\t%s:%d\n\t\t%s\n"
	failOutputWithoutFailLine   = "\t%s:%d\n\t\t%s\n"
	failOutputWithoutLineNumber = "\t\t%s\n"
)

func (logger *errorLoggerImpl) Log(location *location, message string) {
	args := []interface{}{location.Test, location.FileName, location.Line, message}
	if logger.prevTestName != location.Test {
		fmt.Fprintf(logger.writer, failOutput, args...)
	} else if logger.prevTestLine != location.Line {
		fmt.Fprintf(logger.writer, failOutputWithoutFailLine, args[1:]...)
	} else {
		fmt.Fprintf(logger.writer, failOutputWithoutLineNumber, message)
	}
	logger.prevTestName = location.Test
	logger.prevTestLine = location.Line
}
