package assert

import "testing"

func TestThatStringIsEqualToPrintsNothing(t *testing.T) {
	mockT := &mockTestingT{}
	ThatString(mockT, "Alice").IsEqualTo("Alice")
	mockT.assertEmpty(t)
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	mockT := &mockTestingT{}
	ThatString(mockT, "Alice").IsEqualTo("Bob")
	mockT.assertContains(t, "Expected <Bob>, but was <Alice>.\n")
}
