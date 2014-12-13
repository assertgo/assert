package assert

import (
	"bytes"
	"fmt"
	"testing"
)

type mockTestingT struct {
	bytes.Buffer
	t *testing.T
}

func (mockT *mockTestingT) Errorf(format string, args ...interface{}) {
	mockT.WriteString(fmt.Sprintf(format, args...))
}

func (mockT *mockTestingT) assertEmpty() {
	mockT.assertContains("")
}

func (mockT *mockTestingT) assertContains(contents string) {
	ThatString(mockT.t, mockT.String()).IsEqualTo(contents)
}
