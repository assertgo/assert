package assert

type Int64 struct {
	logFacade *logFacade
	actual    int64
}

func (a *Int64) IsZero() *Int64 {
	return a.isTrue(a.actual == 0,
		"Expected zero, but was <%v>.", a.actual)
}

func (a *Int64) IsNonZero() *Int64 {
	return a.isTrue(a.actual != 0,
		"Expected nonzero, but was zero.")
}

func (a *Int64) isTrue(condition bool, format string, args ...interface{}) *Int64 {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
