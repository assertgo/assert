package assert

import (
	"bytes"
	"testing"
)

func TestErrorLoggerSimpleCase(t *testing.T) {
	assert := Setup(t)
	buffer := &bytes.Buffer{}
	logger := &errorLoggerImpl{buffer}
	location := &location{"TestSimple", "simple_test.go", 66}
	message := "Expected KABOOM!"
	logger.Log(location, message)
	assert.ThatString(buffer.String()).IsEqualTo(
		"--- FAIL: TestSimple\n" +
			"\tsimple_test.go:66\n" +
			"\t\tExpected KABOOM!\n")
}
