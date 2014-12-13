package assert

import "testing"

func TestThatStringIsEqualToPrintsNothing(t *testing.T) {
	mockT := &mockTestingT{}
	ThatString(mockT, "Alice").IsEqualTo("Alice")
	assertEmpty(t, mockT)
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	mockT := &mockTestingT{}
	ThatString(mockT, "Alice").IsEqualTo("Bob")
	assertContains(t, mockT, "Expected <Bob>, but was <Alice>.\n")
}
