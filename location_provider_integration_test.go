package assert

import "testing"

func TestLocationProviderViaFramework(t *testing.T) {
	assert, buffer, logger := setupWithLogger(t)
	mockAssert := setupImpl(&testing.T{}, logger)
	mockAssert.ThatBool(false).IsTrue()
	assert.ThatString(buffer.String()).IsEqualTo(
		"\n--- FAIL: TestLocationProviderViaFramework\n" +
			"\tlocation_provider_integration_test.go:8\n" +
			"\t\tExpected <true>, but was <false>.\n")
}
