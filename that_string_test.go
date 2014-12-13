package assert

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestThatStringIsEqualTo(t *testing.T) {
	ThatString("Alice has a cat.").IsEqualTo("Alice has a cat.")
}

func TestThatStringIsEqualToFails(t *testing.T) {
	recoverAndRestore := mockWriter(ioutil.Discard)
	defer recoverAndRestore()
	ThatString("Alice has a cat.").IsEqualTo("Bob has a dog.")
	t.Fail()
}

func TestThatStringIsEqualToPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	assert := func() {
		ThatString(buffer.String()).IsEqualTo("Expected <Alice has a cat.>, but was <Bob has a dog.>.\n")
	}
	recoverAndRestore := mockWriter(buffer, assert)
	defer recoverAndRestore()
	ThatString("Alice has a cat.").IsEqualTo("Bob has a dog.")
}
