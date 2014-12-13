package assert

import (
	"fmt"
	"io"
	"os"
)

var writer io.Writer = os.Stdout

type ThatInt int

func (actual ThatInt) IsZero() ThatInt {
	return actual.IsEqualTo(0)
}

func (actual ThatInt) IsEqualTo(expected int) ThatInt {
	if int(actual) != expected {
		fmt.Fprintf(writer, "Expected <%d>, but was <%d>.\n", expected, actual)
	}
	return actual
}
