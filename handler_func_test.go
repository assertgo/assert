package assert

import (
	"fmt"
	"net/http"
	"testing"
)

func TestThatHandlerFuncWithRequestContainsHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Contains me!")
	}
	assert.ThatHandlerFunc(handler).Contains("Contains me!")
	mockT.HasNoErrors()
}

func TestThatHandlerFuncWithRequestContainsHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Something different.")
	}
	assert.ThatHandlerFunc(handler).Contains("Contains me!")
	mockT.HasErrorMessages(
		"Expected to respond with <Contains me!>, but found <Something different.>.",
	)
}
