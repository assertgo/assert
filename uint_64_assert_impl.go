package assert

type Uint64Assert struct {
	logFacade *logFacade
	actual    uint64
}

func (assert *Uint64Assert) isTrue(condition bool, format string, args ...interface{}) *Uint64Assert {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *Uint64Assert) IsZero() *Uint64Assert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%v>.", assert.actual)
}

func (assert *Uint64Assert) IsNonZero() *Uint64Assert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}
