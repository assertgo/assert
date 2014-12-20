package assert

import (
	"fmt"
	"reflect"
	"testing"
)

type mockTestingT struct {
	t      *testing.T
	errors []string
}

func setupWithMockT(t *testing.T) (assert AssertProvider, mockT *mockTestingT) {
	mockT = newMockT(t)
	assert = Setup(mockT)
	return
}

func newMockT(t *testing.T) *mockTestingT {
	return &mockTestingT{t: t}
}

func (mockT *mockTestingT) Errorf(format string, args ...interface{}) {
	mockT.errors = append(mockT.errors, fmt.Sprintf(format, args...))
}

func (mockT *mockTestingT) HasNoErrors() {
	if len(mockT.errors) > 0 {
		mockT.t.Errorf("Expected to have no errors, but had %d.", len(mockT.errors))
	}
}

func (mockT *mockTestingT) HasErrorMessages(messages ...string) {
	if !reflect.DeepEqual(mockT.errors, messages) {
		mockT.t.Errorf("Expected <%s>, but was <%s>.", messages, mockT.errors)
	}
}
