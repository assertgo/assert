package assert

import "fmt"

func logIfFalse(t TestingT, logger errorLogger, condition bool, format string, args ...interface{}) {
	if !condition {
		location := &location{"TestFooBar", "foo_bar_test.go", 66}
		message := fmt.Sprintf(format, args...)
		logger.Log(location, message)
		t.Fail()
	}
}
