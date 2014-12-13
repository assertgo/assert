package assert

import "testing"

func TestThatStringIsEqualToPrintsNothing(t *testing.T) {
	mockT := &mockTestingT{t: t}
	ThatString(mockT, "Alice").IsEqualTo("Alice")
	mockT.assertEmpty()
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	mockT := &mockTestingT{t: t}
	ThatString(mockT, "Alice").IsEqualTo("Bob")
	mockT.assertContains("Expected <Bob>, but was <Alice>.\n")
}
