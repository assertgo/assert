package assert

import "testing"

func TestThatBoolIsFalseHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatBool(false).IsFalse().IsFalse()
	mockT.HasNoErrors()
}

func TestThatBoolIsFalseHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatBool(true).IsFalse().IsFalse()
	mockT.HasErrorMessages(
		"Expected <false>, but was <true>.",
		"Expected <false>, but was <true>.",
	)
}

func TestThatBoolIsTrueHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatBool(true).IsTrue().IsTrue()
	mockT.HasNoErrors()
}

func TestThatBoolIsTrueHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatBool(false).IsTrue().IsTrue()
	mockT.HasErrorMessages(
		"Expected <true>, but was <false>.",
		"Expected <true>, but was <false>.",
	)
}

func TestThatBoolIsSchrödingersCatHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatBool(false).IsSchrödingersCat()
	assert.ThatBool(true).IsSchrödingersCat()
	mockT.HasErrorMessages(
		"Expected Schrödinger's cat, but was <false>.",
		"Expected Schrödinger's cat, but was <true>.",
	)
}
