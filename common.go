package assert

type T interface {
	Errorf(format string, args ...interface{})
}
