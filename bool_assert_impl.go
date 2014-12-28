package assert

type boolAssertImpl struct {
	logFacade *logFacade
	actual    bool
}

func (assert *boolAssertImpl) isTrue(condition bool, format string, args ...interface{}) *boolAssertImpl {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *boolAssertImpl) IsFalse() BoolAssert {
	return assert.isTrue(assert.actual == false,
		"Expected <false>, but was <%t>.", assert.actual)
}

func (assert *boolAssertImpl) IsTrue() BoolAssert {
	return assert.isTrue(assert.actual == true,
		"Expected <true>, but was <%t>.", assert.actual)
}

func (assert *boolAssertImpl) IsSchrödingersCat() BoolAssert {
	return assert.isTrue(assert.actual == false && assert.actual == true,
		"Expected Schrödinger's cat, but was <%t>.", assert.actual)
}
