package assert

type floatAssertImpl struct {
	logFacade *logFacade
	actual    float64
}

func (assert *floatAssertImpl) isTrue(condition bool, format string, args ...interface{}) *floatAssertImpl {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *floatAssertImpl) IsZero() FloatAssert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%v>.", assert.actual)
}

func (assert *floatAssertImpl) IsNonZero() FloatAssert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}
