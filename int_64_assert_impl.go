package assert

type int64AssertImpl struct {
	logFacade *logFacade
	actual    int64
}

func (assert *int64AssertImpl) isTrue(condition bool, format string, args ...interface{}) *int64AssertImpl {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *int64AssertImpl) IsZero() Int64Assert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%v>.", assert.actual)
}

func (assert *int64AssertImpl) IsNonZero() Int64Assert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}
