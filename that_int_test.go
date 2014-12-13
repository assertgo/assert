package assert

import (
	"bytes"
	"testing"
)

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(t, 0).IsZero()
	writer = writerBackup
	assertEmpty(t, buffer)
}

func TestThatIntChainedPrintsNothing(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(t, 0).IsZero().IsEqualTo(0)
	writer = writerBackup
	assertEmpty(t, buffer)
}

func TestThatIntChainedPrintsAllMessages(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(t, 2).IsZero().IsEqualTo(1)
	writer = writerBackup
	assertContains(t, buffer, "Expected <0>, but was <2>.\nExpected <1>, but was <2>.\n")
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(t, 6).IsZero()
	writer = writerBackup
	assertContains(t, buffer, "Expected <0>, but was <6>.\n")
}

func TestThatIntIsEqualTo(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(t, 1).IsEqualTo(1)
	writer = writerBackup
	assertEmpty(t, buffer)
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	buffer := &bytes.Buffer{}
	writerBackup := writer
	writer = buffer
	ThatInt(t, 1).IsEqualTo(2)
	writer = writerBackup
	assertContains(t, buffer, "Expected <2>, but was <1>.\n")
}

func assertEmpty(t *testing.T, buffer *bytes.Buffer) {
	assertContains(t, buffer, "")
}

func assertContains(t *testing.T, buffer *bytes.Buffer, contents string) {
	ThatString(t, buffer.String()).IsEqualTo(contents)
}
