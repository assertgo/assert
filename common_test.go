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

func (mockT *mockTestingT) HasNoErrors() {
	mockT.HasErrorMessage("")
}

func (mockT *mockTestingT) HasErrorMessage(message string) {
	ThatString(mockT.t, mockT.String()).IsEqualTo(message)
}
