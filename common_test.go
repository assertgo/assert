package assert

import (
	"bytes"
	"fmt"
	"testing"
)

type mockTestingT struct {
	bytes.Buffer
}

func (mockT *mockTestingT) Errorf(format string, args ...interface{}) {
	mockT.WriteString(fmt.Sprintf(format, args...))
}

func assertEmpty(t *testing.T, mockT *mockTestingT) {
	assertContains(t, mockT, "")
}

func assertContains(t *testing.T, mockT *mockTestingT, contents string) {
	ThatString(t, mockT.String()).IsEqualTo(contents)
}
