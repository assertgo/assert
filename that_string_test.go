package assert

import "testing"

func TestThatStringIsEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsEqualTo("Alice").IsEqualTo("Alice")
	mockT.HasNoErrors()
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsEqualTo("Bob").IsEqualTo("Charlie")
	mockT.HasErrorMessages(
		"Expected <Bob>, but was <Alice>.",
		"Expected <Charlie>, but was <Alice>.")
}

func TestThatStringIsNotEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsNotEqualTo("Bob").IsNotEqualTo("Bob")
	mockT.HasNoErrors()
}

func TestThatStringIsNotEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsNotEqualTo("Alice")
	mockT.HasErrorMessages(
		"Expected string not equal to <Alice>, but was equal.")
}
