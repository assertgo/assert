package assert

type FloatAssert struct {
	logFacade *logFacade
	actual    float64
}

func (assert *FloatAssert) isTrue(condition bool, format string, args ...interface{}) *FloatAssert {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *FloatAssert) IsZero() *FloatAssert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%v>.", assert.actual)
}

func (assert *FloatAssert) IsNonZero() *FloatAssert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}
