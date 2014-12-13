package assert

import (
	"bytes"
	"io"
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

func TestThatIntChainedPrintsAllMessages(t *testing.T) {
	buffer := &bytes.Buffer{}
	writer = buffer
	ThatInt(2).IsZero().IsEqualTo(1)
	assertContains(buffer, "Expected <0>, but was <2>.\nExpected <1>, but was <2>.\n")
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	writer = buffer
	ThatInt(6).IsZero()
	ThatString(buffer.String()).IsEqualTo("Expected <0>, but was <6>.\n")
}

func TestThatIntIsEqualTo(t *testing.T) {
	ThatInt(1).IsEqualTo(1)
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	writer = buffer
	ThatInt(1).IsEqualTo(2)
	ThatString(buffer.String()).IsEqualTo("Expected <2>, but was <1>.\n")
}

func mockWriter(writerDouble io.Writer, asserts ...func()) func() {
	writerBackup := writer
	writer = writerDouble
	return func() {
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

func assertContains(buffer *bytes.Buffer, contents string) {
	ThatString(buffer.String()).IsEqualTo(contents)
}
