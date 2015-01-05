package assert

type Float struct {
	logFacade *logFacade
	actual    float64
}

func (a *Float) IsZero() *Float {
	return a.isTrue(a.actual == 0,
		"Expected zero, but was <%v>.", a.actual)
}

func (a *Float) IsNonZero() *Float {
	return a.isTrue(a.actual != 0,
		"Expected nonzero, but was zero.")
}

func (a *Float) isTrue(condition bool, format string, args ...interface{}) *Float {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
