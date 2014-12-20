package assert

import "fmt"

func logIfFalse(logFacade *logFacade, condition bool, format string, args ...interface{}) {
	if !condition {
		location := &location{"TestFooBar", "foo_bar_test.go", 66}
		message := fmt.Sprintf(format, args...)
		logFacade.Log(location, message)
	}
}
