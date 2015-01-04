package assert

func (a *Int64) isTrue(condition bool, format string, args ...interface{}) *Int64 {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
