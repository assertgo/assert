package assert

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestThatStringIsEqualTo(t *testing.T) {
	ThatString("Alice").IsEqualTo("Alice")
}

func TestThatStringIsEqualToFails(t *testing.T) {
	recoverAndRestore := mockWriter(ioutil.Discard)
	defer recoverAndRestore()
	ThatString("Alice").IsEqualTo("Bob")
	t.Fail()
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
