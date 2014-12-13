package assert

import (
	"fmt"
	"io"
	"os"
	"testing"
)

var writer io.Writer = os.Stdout

func ThatInt(t *testing.T, actual int) *IntAssert {
	return &IntAssert{t, actual}
}

type IntAssert struct {
	t      *testing.T
	actual int
}

func (assert *IntAssert) IsZero() *IntAssert {
	return assert.IsEqualTo(0)
}

func (assert *IntAssert) IsEqualTo(expected int) *IntAssert {
	if assert.actual != expected {
		fmt.Fprintf(writer, "Expected <%d>, but was <%d>.\n", expected, assert.actual)
	}
	return assert
}
