package assert

type TestingT interface {
	Errorf(format string, args ...interface{})
}
