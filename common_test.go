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

func (mockT *mockTestingT) assertEmpty(t *testing.T) {
	mockT.assertContains(t, "")
}

func (mockT *mockTestingT) assertContains(t *testing.T, contents string) {
	ThatString(t, mockT.String()).IsEqualTo(contents)
}
