package assert

import "testing"

func TestThatComplexIsZeroHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatComplex(0i).IsZero().IsZero()
	mockT.HasNoErrors()
}

func TestThatComplexIsZeroHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatComplex(-1).IsZero()
	assert.ThatComplex(1i).IsZero()
	mockT.HasErrorMessages(
		"Expected zero, but was <(-1+0i)>.",
		"Expected zero, but was <(0+1i)>.",
	)
}

func TestThatComplexIsNonZeroHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatComplex(-1).IsNonZero()
	assert.ThatComplex(1i).IsNonZero()
	mockT.HasNoErrors()
}

func TestThatComplexIsNonZeroHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatComplex(complex(0, 0)).IsNonZero().IsNonZero()
	mockT.HasErrorMessages(
		"Expected nonzero, but was zero.",
		"Expected nonzero, but was zero.",
	)
}
