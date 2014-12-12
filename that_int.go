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
		fmt.Fprintln(writer, "Expected <0>, but was <1>.")
		panic("that int should be zero")
	}
}
