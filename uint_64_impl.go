package assert

func (a *Uint64) isTrue(condition bool, format string, args ...interface{}) *Uint64 {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}
