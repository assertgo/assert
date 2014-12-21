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
	fmt.Fprintf(logger.writer, "--- FAIL: TestSimple\n\tsimple_test.go:66\n\t\tExpected KABOOM!\n")
}
