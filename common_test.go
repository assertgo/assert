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

func assertEmpty(t *testing.T, buffer *BufferT) {
	assertContains(t, buffer, "")
}

func assertContains(t *testing.T, buffer *BufferT, contents string) {
	ThatString(t, buffer.b.String()).IsEqualTo(contents)
}
