package assert

func (a *Bool) isTrue(condition bool, format string, args ...interface{}) *Bool {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
