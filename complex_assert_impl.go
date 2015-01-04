package assert

type complexAssertImpl struct {
	logFacade *logFacade
	actual    complex128
}

func (assert *complexAssertImpl) isTrue(condition bool, format string, args ...interface{}) *complexAssertImpl {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *complexAssertImpl) IsZero() ComplexAssert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%v>.", assert.actual)
}

func (assert *complexAssertImpl) IsNonZero() ComplexAssert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}
