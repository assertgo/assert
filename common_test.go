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

func newMockT(t *testing.T) *mockTestingT {
	return &mockTestingT{t: t}
}

func (mockT *mockTestingT) Errorf(format string, args ...interface{}) {
	mockT.WriteString(fmt.Sprintf(format, args...))
}

func (mockT *mockTestingT) HasNoErrors() {
	mockT.HasErrorMessage("")
}

func (mockT *mockTestingT) HasErrorMessage(message string) {
	if mockT.String() != message {
		mockT.t.Errorf("Expected <%s>, but was <%s>.", message, mockT.String())
	}
}
