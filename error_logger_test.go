package assert

import (
	"bytes"
	"testing"
)

func TestErrorLoggerSimpleCase(t *testing.T) {
	assert, buffer, logger := setupWithLogger(t)
	location := &location{"TestSimple", "simple_test.go", 66}
	message := "Expected KABOOM!"
	logger.Log(location, message)
	assert.ThatString(buffer.String()).IsEqualTo(
		"--- FAIL: TestSimple\n" +
			"\tsimple_test.go:66\n" +
			"\t\tExpected KABOOM!\n")
}

func setupWithLogger(t *testing.T) (assert AssertProvider, buffer *bytes.Buffer, logger errorLogger) {
	assert = Setup(t)
	buffer = &bytes.Buffer{}
	logger = &errorLoggerImpl{buffer}
	return
}
