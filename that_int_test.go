package assert

import (
	"bytes"
	"fmt"
	"testing"
)

type BufferT struct {
	b bytes.Buffer
}

func (buffer *BufferT) Errorf(format string, args ...interface{}) {
	buffer.b.WriteString(fmt.Sprintf(format, args...))
}

func TestThatIntIsZeroPrintsNothing(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 0).IsZero()
	assertEmpty(t, buffer)
}

func TestThatIntChainedPrintsNothing(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 0).IsZero().IsEqualTo(0)
	assertEmpty(t, buffer)
}

func TestThatIntChainedPrintsAllMessages(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 2).IsZero().IsEqualTo(1)
	assertContains(t, buffer, "Expected <0>, but was <2>.\nExpected <1>, but was <2>.\n")
}

func TestThatIntIsZeroPrintsMessage(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 6).IsZero()
	assertContains(t, buffer, "Expected <0>, but was <6>.\n")
}

func TestThatIntIsEqualTo(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 1).IsEqualTo(1)
	assertEmpty(t, buffer)
}

func TestThatIntIsEqualToPrintsMessage(t *testing.T) {
	buffer := &BufferT{}
	ThatInt(buffer, 1).IsEqualTo(2)
	assertContains(t, buffer, "Expected <2>, but was <1>.\n")
}

func assertEmpty(t *testing.T, buffer *BufferT) {
	assertContains(t, buffer, "")
}

func assertContains(t *testing.T, buffer *BufferT, contents string) {
	ThatString(t, buffer.b.String()).IsEqualTo(contents)
}
