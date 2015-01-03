package assert

import "testing"

func TestThatInt64IsZeroHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatInt64(0).IsZero().IsZero()
	mockT.HasNoErrors()
}

func TestThatInt64IsZeroHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatInt64(-1).IsZero()
	assert.ThatInt64(1).IsZero()
	mockT.HasErrorMessages(
		"Expected zero, but was <-1>.",
		"Expected zero, but was <1>.",
	)
}

func TestThatInt64IsNonZeroHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatInt64(-1).IsNonZero()
	assert.ThatInt64(1).IsNonZero()
	mockT.HasNoErrors()
}

func TestThatInt64IsNonZeroHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatInt64(0).IsNonZero().IsNonZero()
	mockT.HasErrorMessages(
		"Expected nonzero, but was zero.",
		"Expected nonzero, but was zero.",
	)
}
