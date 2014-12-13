package assert

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestThatIntIsZero(t *testing.T) {
	ThatInt(0).IsZero()
}

func TestThatIntIsZeroFails(t *testing.T) {
	recoverAndRestore := mockWriter(ioutil.Discard)
	defer recoverAndRestore()
	ThatInt(1).IsZero()
	t.Fail()
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	assert := func() {
		if buffer.String() != "Expected <0>, but was <1>.\n" {
			t.Fail()
		}
	}
	recoverAndRestore := mockWriter(buffer, assert)
	defer recoverAndRestore()
	ThatInt(1).IsZero()
}

func mockWriter(writerDouble io.Writer, asserts ...func()) func() {
	writerBackup := writer
	writer = writerDouble
	return func() {
		recover()
		callAsserts(asserts...)
		writer = writerBackup
	}
}

func callAsserts(asserts ...func()) {
	for _, assert := range asserts {
		assert()
	}
}
