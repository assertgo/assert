package assert

import "testing"

func TestThatBoolIsFalseHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	assert := Setup(mockT)
	assert.ThatBool(false).IsFalse().IsFalse()
	mockT.HasNoErrors()
}

func TestThatBoolIsFalseHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	assert := Setup(mockT)
	assert.ThatBool(true).IsFalse().IsFalse()
	mockT.HasErrorMessages(
		"Expected <false>, but was <true>.",
		"Expected <false>, but was <true>.",
	)
}

func TestThatBoolIsTrueHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	assert := Setup(mockT)
	assert.ThatBool(true).IsTrue().IsTrue()
	mockT.HasNoErrors()
}

func TestThatBoolIsTrueHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	assert := Setup(mockT)
	assert.ThatBool(false).IsTrue().IsTrue()
	mockT.HasErrorMessages(
		"Expected <true>, but was <false>.",
		"Expected <true>, but was <false>.",
	)
}
