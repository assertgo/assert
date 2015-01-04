package assert

type Int64Assert struct {
	logFacade *logFacade
	actual    int64
}

func (assert *Int64Assert) isTrue(condition bool, format string, args ...interface{}) *Int64Assert {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *Int64Assert) IsZero() *Int64Assert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%v>.", assert.actual)
}

func (assert *Int64Assert) IsNonZero() *Int64Assert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}
