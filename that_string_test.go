package assert

import (
	"bytes"
	"testing"
)

func TestThatStringIsEqualTo(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatString("Alice").IsEqualTo("Alice")
	writer = writerBackup
	assertEmpty(buffer)
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatString("Alice").IsEqualTo("Bob")
	writer = writerBackup
	assertContains(buffer, "Expected <Bob>, but was <Alice>.\n")
}
