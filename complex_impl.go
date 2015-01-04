package assert

func (a *Complex) isTrue(condition bool, format string, args ...interface{}) *Complex {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
