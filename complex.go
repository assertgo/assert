package assert

type Complex struct {
	logFacade *logFacade
	actual    complex128
}

func (a *Complex) IsZero() *Complex {
	return a.isTrue(a.actual == 0,
		"Expected zero, but was <%v>.", a.actual)
}

func (a *Complex) IsNonZero() *Complex {
	return a.isTrue(a.actual != 0,
		"Expected nonzero, but was zero.")
}

func (a *Complex) IsEqualTo(expected complex128) *Complex {
	return a.isTrue(a.actual == expected,
		"Expected <%v>, but was <%v>.", expected, a.actual)
}

func (a *Complex) IsNotEqualTo(unexpected complex128) *Complex {
	return a.isTrue(a.actual != unexpected,
		"Expected not equal to <%v>, but was.", unexpected)
}

func (a *Complex) isTrue(condition bool, format string, args ...interface{}) *Complex {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
