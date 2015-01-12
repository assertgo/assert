package assert

import (
	"net/http"
	"net/http/httptest"
)

type HandlerFunc struct {
	logFacade *logFacade
	actual    http.HandlerFunc
}

func (a *HandlerFunc) Contains(expected string) *HandlerFunc {
	w := httptest.NewRecorder()
	a.actual(w, nil)
	body := w.Body.String()
	return a.isTrue(body == expected,
		"Expected to respond with <%s>, but found <%s>.", expected, body)
}

func (a *HandlerFunc) HasCode(expected int) *HandlerFunc {
	w := httptest.NewRecorder()
	a.actual(w, nil)
	code := w.Code
	return a.isTrue(code == expected,
		"Expected to have code <%d>, but found <%d>.", expected, code)
}

func (a *HandlerFunc) isTrue(condition bool, format string, args ...interface{}) *HandlerFunc {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
