package assert

import (
	"bytes"
	"testing"
)

func TestThatStringIsEqualTo(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatString(t, "Alice").IsEqualTo("Alice")
	writer = writerBackup
	assertEmpty(t, buffer)
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatString(t, "Alice").IsEqualTo("Bob")
	writer = writerBackup
	assertContains(t, buffer, "Expected <Bob>, but was <Alice>.\n")
}
