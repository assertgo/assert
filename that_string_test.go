package assert

import (
	"bytes"
	"testing"
)

func TestThatStringIsEqualTo(t *testing.T) {
	ThatString("Alice").IsEqualTo("Alice")
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	assert := func() {
		ThatString(buffer.String()).IsEqualTo("Expected <Bob>, but was <Alice>.\n")
	}
	recoverAndRestore := mockWriter(buffer, assert)
	defer recoverAndRestore()
	ThatString("Alice").IsEqualTo("Bob")
}
