package assert

import (
	"fmt"
	"net/http"
	"testing"
)

func TestThatHandlerFuncContainsHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Contains me!")
	}
	assert.ThatHandlerFunc(handler).Contains("Contains me!")
	mockT.HasNoErrors()
}

func TestThatHandlerFuncContainsHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Something different.")
	}
	assert.ThatHandlerFunc(handler).Contains("Contains me!")
	mockT.HasErrorMessages(
		"Expected to respond with <Contains me!>, but found <Something different.>.",
	)
}

func TestThatHandlerFuncHasCodeHasNoErrors(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatHandlerFunc(func(w http.ResponseWriter, r *http.Request) {}).HasCode(http.StatusOK)
	assert.ThatHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}).HasCode(http.StatusNotFound)
	mockT.HasNoErrors()
}

func TestThatHandlerFuncHasCodeHasErrorMessages(t *testing.T) {
	assert, mockT := setupWithMockT(t)
	assert.ThatHandlerFunc(func(w http.ResponseWriter, r *http.Request) {}).
		HasCode(http.StatusNotModified).
		HasCode(http.StatusNotFound)
	mockT.HasErrorMessages(
		"Expected to have code <304>, but found <200>.",
		"Expected to have code <404>, but found <200>.",
	)
}
