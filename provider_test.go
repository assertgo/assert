package assert

import (
	"reflect"
	"testing"
)

type mockTestingT struct {
	t      *testing.T
	failed bool
	errors []string
}

func setupWithMockT(t *testing.T) (provider *Provider, mockT *mockTestingT) {
	mockT = &mockTestingT{t: t}
	provider = setupImpl(mockT, mockT)
	return
}

func (mockT *mockTestingT) Log(location *location, message string) {
	mockT.errors = append(mockT.errors, message)
}

func (mockT *mockTestingT) Fail() {
	mockT.failed = true
}

func (mockT *mockTestingT) HasNoErrors() {
	if mockT.failed || len(mockT.errors) > 0 {
		mockT.t.Errorf("Expected to have no errors, but had <%s>.", mockT.errors)
	}
}

func (mockT *mockTestingT) HasErrorMessages(messages ...string) {
	if !mockT.failed || !reflect.DeepEqual(mockT.errors, messages) {
		mockT.t.Errorf("Expected <%s>, but was <%s>.", messages, mockT.errors)
	}
}
