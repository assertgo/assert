package assert

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	buffer := &bytes.Buffer{}
	writer = buffer
	ThatInt(0).IsZero()
	assertEmpty(buffer)
}

func TestThatIntChainedPrintsNothing(t *testing.T) {
	buffer := &bytes.Buffer{}
	writer = buffer
	ThatInt(0).IsZero().IsEqualTo(0)
	assertEmpty(buffer)
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
		ThatString(buffer.String()).IsEqualTo("Expected <0>, but was <6>.\n")
	}
	recoverAndRestore := mockWriter(buffer, assert)
	defer recoverAndRestore()
	ThatInt(6).IsZero()
}

func TestThatIntIsEqualTo(t *testing.T) {
	ThatInt(1).IsEqualTo(1)
}

func TestThatIntIsEqualToFails(t *testing.T) {
	recoverAndRestore := mockWriter(ioutil.Discard)
	defer recoverAndRestore()
	ThatInt(1).IsEqualTo(2)
	t.Fail()
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	assert := func() {
		ThatString(buffer.String()).IsEqualTo("Expected <2>, but was <1>.\n")
	}
	recoverAndRestore := mockWriter(buffer, assert)
	defer recoverAndRestore()
	ThatInt(1).IsEqualTo(2)
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

func assertEmpty(buffer *bytes.Buffer) {
	ThatString(buffer.String()).IsEqualTo("")
}
