package assert

type BoolAssert struct {
	logFacade *logFacade
	actual    bool
}

func (assert *BoolAssert) isTrue(condition bool, format string, args ...interface{}) *BoolAssert {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *BoolAssert) IsFalse() *BoolAssert {
	return assert.isTrue(assert.actual == false,
		"Expected <false>, but was <%t>.", assert.actual)
}

func (assert *BoolAssert) IsTrue() *BoolAssert {
	return assert.isTrue(assert.actual == true,
		"Expected <true>, but was <%t>.", assert.actual)
}

func (assert *BoolAssert) IsSchrödingersCat() *BoolAssert {
	return assert.isTrue(assert.actual == false && assert.actual == true,
		"Expected Schrödinger's cat, but was <%t>.", assert.actual)
}
