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

func TestThatComplexIsEqualToHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatComplex(1i).IsEqualTo(complex(0, 1)).IsEqualTo(0 + 1i)
	mockT.HasNoErrors()
}

func TestThatComplexIsEqualToHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatComplex(1i).IsEqualTo(1).IsEqualTo(2i)
	mockT.HasErrorMessages(
		"Expected <(1+0i)>, but was <(0+1i)>.",
		"Expected <(0+2i)>, but was <(0+1i)>.",
	)
}

func TestThatComplexIsNotEqualToHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatComplex(1i).IsNotEqualTo(1).IsNotEqualTo(2i)
	mockT.HasNoErrors()
}

func TestThatComplexIsNotEqualToHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatComplex(1i).IsNotEqualTo(complex(0, 1))
	assert.ThatComplex(0.3 - 0.3i).IsNotEqualTo(complex(0.3, -0.3))
	mockT.HasErrorMessages(
		"Expected not equal to <(0+1i)>, but was.",
		"Expected not equal to <(0.3-0.3i)>, but was.",
	)
}
