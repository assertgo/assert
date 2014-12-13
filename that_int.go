package assert

import (
	"fmt"
	"io"
	"os"
)

var writer io.Writer = os.Stdout

type ThatInt int

func (thatInt ThatInt) IsZero() ThatInt {
	if thatInt != 0 {
		fmt.Fprintf(writer, "Expected <0>, but was <%d>.\n", thatInt)
		panic("Expected int to be zero but wasn't.")
	}
	return thatInt
}

func (thatInt ThatInt) IsEqualTo(comparedInt int) ThatInt {
	if int(thatInt) != comparedInt {
		fmt.Fprintf(writer, "Expected <%d>, but was <%d>.\n", thatInt, comparedInt)
		panic("Expected int to be equal but wasn't.")
	}
	return thatInt
}
