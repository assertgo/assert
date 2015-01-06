package assert

import (
	"bytes"
	"testing"
)

func TestErrorLoggerSimpleCase(t *testing.T) {
	assert, buffer, logger := setupWithLogger(t)
	logger.Log(&location{"TestSimple", "simple_test.go", 66}, "Expected KABOOM!")
	assert.ThatString(buffer.String()).IsEqualTo(
		"\n--- FAIL: TestSimple\n" +
			"\tsimple_test.go:66\n" +
			"\t\tExpected KABOOM!\n")
}

func TestErrorLoggerSecondError(t *testing.T) {
	assert, buffer, logger := setupWithLogger(t)
	logger.Log(&location{"TestSimple", "simple_test.go", 66}, "Expected KABOOM!")
	logger.Log(&location{"TestSimple", "simple_test.go", 67}, "Expected BAZOOKA?")
	assert.ThatString(buffer.String()).IsEqualTo(
		"\n--- FAIL: TestSimple\n" +
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
		"\n--- FAIL: TestSimple\n" +
			"\tsimple_test.go:66\n" +
			"\t\tExpected KABOOM!\n" +
			"\n--- FAIL: TestNotSoSimple\n" +
			"\tsimple_test.go:70\n" +
			"\t\tExpected BAZOOKA?\n")
}

func TestErrorLoggerSecondErrorOnTheSameLine(t *testing.T) {
	assert, buffer, logger := setupWithLogger(t)
	logger.Log(&location{"TestChain", "chain_test.go", 42}, "Expected CHAIN!")
	logger.Log(&location{"TestChain", "chain_test.go", 42}, "Expected SAW?")
	assert.ThatString(buffer.String()).IsEqualTo(
		"\n--- FAIL: TestChain\n" +
			"\tchain_test.go:42\n" +
			"\t\tExpected CHAIN!\n" +
			"\t\tExpected SAW?\n")
}

func TestErrorLoggerDifferentFileSameLine(t *testing.T) {
	assert, buffer, logger := setupWithLogger(t)
	logger.Log(&location{"TestFirst", "first_test.go", 99}, "Expected First!")
	logger.Log(&location{"TestSecond", "second_test.go", 99}, "Expected Second?")
	assert.ThatString(buffer.String()).IsEqualTo(
		"\n--- FAIL: TestFirst\n" +
			"\tfirst_test.go:99\n" +
			"\t\tExpected First!\n" +
			"\n--- FAIL: TestSecond\n" +
			"\tsecond_test.go:99\n" +
			"\t\tExpected Second?\n")
}

func setupWithLogger(t *testing.T) (provider *Provider, buffer *bytes.Buffer, logger errorLogger) {
	provider = New(t)
	buffer = &bytes.Buffer{}
	logger = &errorLoggerImpl{writer: buffer}
	return
}
