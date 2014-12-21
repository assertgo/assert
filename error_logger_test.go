package assert

import (
	"bytes"
	"testing"
)

func TestErrorLoggerSimpleCase(t *testing.T) {
	assert, buffer, logger := setupWithLogger(t)
	logger.Log(&location{"TestSimple", "simple_test.go", 66}, "Expected KABOOM!")
	assert.ThatString(buffer.String()).IsEqualTo(
		"--- FAIL: TestSimple\n" +
			"\tsimple_test.go:66\n" +
			"\t\tExpected KABOOM!\n")
}

func TestErrorLoggerSecondError(t *testing.T) {
	assert, buffer, logger := setupWithLogger(t)
	logger.Log(&location{"TestSimple", "simple_test.go", 66}, "Expected KABOOM!")
	logger.Log(&location{"TestSimple", "simple_test.go", 67}, "Expected BAZOOKA?")
	assert.ThatString(buffer.String()).IsEqualTo(
		"--- FAIL: TestSimple\n" +
			"\tsimple_test.go:66\n" +
			"\t\tExpected KABOOM!\n" +
			"\tsimple_test.go:67\n" +
			"\t\tExpected BAZOOKA?\n")
}

func TestErrorLoggerTwoTestsFailing(t *testing.T) {
	assert, buffer, logger := setupWithLogger(t)
	logger.Log(&location{"TestSimple", "simple_test.go", 66}, "Expected KABOOM!")
	logger.Log(&location{"TestNotSoSimple", "simple_test.go", 70}, "Expected BAZOOKA?")
	assert.ThatString(buffer.String()).IsEqualTo(
		"--- FAIL: TestSimple\n" +
			"\tsimple_test.go:66\n" +
			"\t\tExpected KABOOM!\n" +
			"--- FAIL: TestNotSoSimple\n" +
			"\tsimple_test.go:70\n" +
			"\t\tExpected BAZOOKA?\n")
}

func setupWithLogger(t *testing.T) (assert AssertProvider, buffer *bytes.Buffer, logger errorLogger) {
	assert = Setup(t)
	buffer = &bytes.Buffer{}
	logger = &errorLoggerImpl{writer: buffer}
	return
}
