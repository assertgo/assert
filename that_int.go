package assert

import (
	"fmt"
	"io"
	"os"
)

var writer io.Writer = os.Stdout

type ThatInt int

func (thatInt ThatInt) IsZero() ThatInt {
	return thatInt.IsEqualTo(0)
}

func (thatInt ThatInt) IsEqualTo(comparedInt int) ThatInt {
	if int(thatInt) != comparedInt {
		fmt.Fprintf(writer, "Expected <%d>, but was <%d>.\n", comparedInt, thatInt)
		panic("Expected int to be equal but wasn't.")
	}
	return thatInt
}
