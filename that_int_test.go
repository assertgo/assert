package assert

import (
	"bytes"
	"testing"
)

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(0).IsZero()
	writer = writerBackup
	assertEmpty(buffer)
}

func TestThatIntChainedPrintsNothing(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(0).IsZero().IsEqualTo(0)
	writer = writerBackup
	assertEmpty(buffer)
}

func TestThatIntChainedPrintsAllMessages(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(2).IsZero().IsEqualTo(1)
	writer = writerBackup
	assertContains(buffer, "Expected <0>, but was <2>.\nExpected <1>, but was <2>.\n")
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(6).IsZero()
	writer = writerBackup
	assertContains(buffer, "Expected <0>, but was <6>.\n")
}

func TestThatIntIsEqualTo(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(1).IsEqualTo(1)
	writer = writerBackup
	assertEmpty(buffer)
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(1).IsEqualTo(2)
	writer = writerBackup
	assertContains(buffer, "Expected <2>, but was <1>.\n")
}

func assertEmpty(buffer *bytes.Buffer) {
	assertContains(buffer, "")
}

func assertContains(buffer *bytes.Buffer, contents string) {
	ThatString(buffer.String()).IsEqualTo(contents)
}
