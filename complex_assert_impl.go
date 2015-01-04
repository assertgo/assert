package assert

type ComplexAssert struct {
	logFacade *logFacade
	actual    complex128
}

func (assert *ComplexAssert) isTrue(condition bool, format string, args ...interface{}) *ComplexAssert {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *ComplexAssert) IsZero() *ComplexAssert {
	return assert.isTrue(assert.actual == 0,
		"Expected zero, but was <%v>.", assert.actual)
}

func (assert *ComplexAssert) IsNonZero() *ComplexAssert {
	return assert.isTrue(assert.actual != 0,
		"Expected nonzero, but was zero.")
}
