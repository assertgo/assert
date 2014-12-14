package assert

import "testing"

func TestThatStringIsEqualToPrintsNothing(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsEqualTo("Alice")
	mockT.HasNoErrors()
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	mockT := newMockT(t)
	ThatString(mockT, "Alice").IsEqualTo("Bob")
	mockT.HasErrorMessages("Expected <Bob>, but was <Alice>.\n")
}
