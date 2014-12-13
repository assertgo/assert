package assert

import (
	"fmt"
	"testing"
)

func ThatString(t *testing.T, actual string) *StringAssert {
	return &StringAssert{t, actual}
}

type StringAssert struct {
	t      *testing.T
	actual string
}

func (assert *StringAssert) IsEqualTo(expected string) *StringAssert {
	if assert.actual != expected {
		fmt.Fprintf(writer, "Expected <%s>, but was <%s>.\n", expected, assert.actual)
	}
	return assert
}
