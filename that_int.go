package assert

import (
	"fmt"
	"io"
	"os"
)

var writer io.Writer = os.Stdout

type ThatInt int

func (thatInt ThatInt) IsZero() {
	if thatInt != 0 {
		fmt.Fprintf(writer, "Expected <0>, but was <%d>.\n", thatInt)
		panic("that int should be zero")
	}
}

func (thatInt ThatInt) IsEqualTo(comparedInt int) {
	if int(thatInt) != comparedInt {
		fmt.Fprintf(writer, "Expected <%d>, but was <%d>.\n", thatInt, comparedInt)
		panic("that int should be equal to compared int")
	}
}
