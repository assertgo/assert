package assert

func logIfFalse(t TestingT, condition bool, format string, args ...interface{}) {
	if !condition {
		t.Errorf(format, args...)
	}
}
