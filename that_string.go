package assert

import "fmt"

type ThatString string

func (actual ThatString) IsEqualTo(expected string) {
	if string(actual) != expected {
		message := fmt.Sprintf("Expected <%s>, but was <%s>.\n", expected, actual)
		fmt.Fprint(writer, message)
	}
}
