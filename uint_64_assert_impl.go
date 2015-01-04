package assert

type uint64AssertImpl struct {
	logFacade *logFacade
	actual    uint64
}

func (assert *uint64AssertImpl) isTrue(condition bool, format string, args ...interface{}) *uint64AssertImpl {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *uint64AssertImpl) IsZero() Uint64Assert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%v>.", assert.actual)
}

func (assert *uint64AssertImpl) IsNonZero() Uint64Assert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}
