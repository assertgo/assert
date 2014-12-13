package assert

import "testing"

func TestThatStringIsEqualTo(t *testing.T) {
	buffer := &BufferT{}
	ThatString(buffer, "Alice").IsEqualTo("Alice")
	assertEmpty(t, buffer)
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	buffer := &BufferT{}
	ThatString(buffer, "Alice").IsEqualTo("Bob")
	assertContains(t, buffer, "Expected <Bob>, but was <Alice>.\n")
}
