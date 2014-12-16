package assert

import "testing"

func TestThatBoolIsFalseHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatBool(mockT, false).IsFalse().IsFalse()
	mockT.HasNoErrors()
}

func TestThatBoolIsFalseHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatBool(mockT, true).IsFalse().IsFalse()
	mockT.HasErrorMessages(
		"Expected <false>, but was <true>.",
		"Expected <false>, but was <true>.",
	)
}

func TestThatBoolIsTrueHasNoErrors(t *testing.T) {
	mockT := newMockT(t)
	ThatBool(mockT, true).IsTrue().IsTrue()
	mockT.HasNoErrors()
}

func TestThatBoolIsTrueHasErrorMessages(t *testing.T) {
	mockT := newMockT(t)
	ThatBool(mockT, false).IsTrue().IsTrue()
	mockT.HasErrorMessages(
		"Expected <true>, but was <false>.",
		"Expected <true>, but was <false>.",
	)
}
