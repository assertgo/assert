package assert

func (a *Float) isTrue(condition bool, format string, args ...interface{}) *Float {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
