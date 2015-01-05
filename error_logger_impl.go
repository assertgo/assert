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
	formatHeaderFull  = "\n--- FAIL: %s\n\t%s:%d\n"
	formatHeaderShort = "\t%s:%d\n"
	formatMessage     = "\t\t%s\n"
)

func (logger *errorLoggerImpl) Log(location *location, message string) {
	if logger.prevTestName != location.Test {
		fmt.Fprintf(logger.writer, formatHeaderFull, location.Test, location.FileName, location.Line)
	} else if logger.prevTestLine != location.Line {
		fmt.Fprintf(logger.writer, formatHeaderShort, location.FileName, location.Line)
	}
	fmt.Fprintf(logger.writer, formatMessage, message)
	logger.prevTestName = location.Test
	logger.prevTestLine = location.Line
}
