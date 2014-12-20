package assert

import "testing"

func TestCanCreateAssert(t *testing.T) {
	mockT := newMockT(t)
	assert := Setup(mockT)
	assert.ThatInt(0).IsZero()
	assert.ThatString("").IsEmpty()
	assert.ThatBool(false).IsFalse()
	mockT.HasNoErrors()
}

func TestCanUseThatFunctionsWithoutErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 0).IsZero()
	ThatString(mockT, "").IsEmpty()
	ThatBool(mockT, false).IsFalse()
	mockT.HasNoErrors()
}

func TestCanUseThatFunctionsWithErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatInt(mockT, 1).IsZero()
	ThatString(mockT, "a").IsEmpty()
	ThatBool(mockT, true).IsFalse()
	mockT.HasErrorMessages(
		"Expected zero, but was <1>.",
		"Expected string to be empty, but was <a>.",
		"Expected <false>, but was <true>.",
	)
}
