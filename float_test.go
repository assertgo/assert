package assert

import "testing"

func TestThatFloatIsZeroHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatFloat(0.0).IsZero().IsZero()
	mockT.HasNoErrors()
}

func TestThatFloatIsZeroHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatFloat(-1).IsZero()
	assert.ThatFloat(0.1).IsZero()
	mockT.HasErrorMessages(
		"Expected zero, but was <-1>.",
		"Expected zero, but was <0.1>.",
	)
}

func TestThatFloatIsNonZeroHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatFloat(-0.1).IsNonZero()
	assert.ThatFloat(1).IsNonZero()
	mockT.HasNoErrors()
}

func TestThatFloatIsNonZeroHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatFloat(0.0).IsNonZero().IsNonZero()
	mockT.HasErrorMessages(
		"Expected nonzero, but was zero.",
		"Expected nonzero, but was zero.",
	)
}
