package assert

import (
	"bytes"
	"testing"
)

func TestLocationProviderViaFramework(t *testing.T) {
	assert := Setup(t)
	mockT := &mockTestingT{t: t}
	buffer := &bytes.Buffer{}
	logger := &errorLoggerImpl{writer: buffer}
	mockAssert := setupImpl(mockT, logger)
	mockAssert.ThatBool(false).IsTrue()
	assert.ThatString(buffer.String()).IsEqualTo(
		"\n--- FAIL: TestLocationProviderViaFramework\n" +
			"\tlocation_provider_integration_test.go:14\n" +
			"\t\tExpected <true>, but was <false>.\n")
}
