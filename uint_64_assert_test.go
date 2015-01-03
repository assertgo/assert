package assert

import "testing"

func TestThatUint64IsZeroHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatUint64(0).IsZero().IsZero()
	mockT.HasNoErrors()
}

func TestThatUint64IsZeroHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatUint64(1).IsZero()
	mockT.HasErrorMessages(
		"Expected zero, but was <1>.",
	)
}

func TestThatUint64IsNonZeroHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatUint64(1).IsNonZero()
	mockT.HasNoErrors()
}

func TestThatUint64IsNonZeroHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatUint64(0).IsNonZero().IsNonZero()
	mockT.HasErrorMessages(
		"Expected nonzero, but was zero.",
		"Expected nonzero, but was zero.",
	)
}
