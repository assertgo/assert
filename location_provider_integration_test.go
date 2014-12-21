package assert

import (
	"bytes"
	"testing"
)

func TestLocationProviderViaFramework(t *testing.T) {
	assert := Setup(t)
	buffer := &bytes.Buffer{}
	logger := &errorLoggerImpl{writer: buffer}
	mockAssert := setupImpl(&testing.T{}, logger)
	mockAssert.ThatBool(false).IsTrue()
	assert.ThatString(buffer.String()).IsEqualTo(
		"\n--- FAIL: TestLocationProviderViaFramework\n" +
			"\tlocation_provider_integration_test.go:13\n" +
			"\t\tExpected <true>, but was <false>.\n")
}
