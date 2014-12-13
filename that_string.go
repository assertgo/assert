package assert

import "fmt"

type ThatString string

func (thatString ThatString) IsEqualTo(comparedString string) {
	if string(thatString) != comparedString {
		fmt.Fprintf(writer, "Expected <%s>, but was <%s>.\n", thatString, comparedString)
		panic("")
	}
}
