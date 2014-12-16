package assert

type boolAssertImpl struct {
	t      TestingT
	actual bool
}

func (assert *boolAssertImpl) isTrue(condition bool, format string, args ...interface{}) *boolAssertImpl {
	logIfFalse(assert.t, condition, format, args...)
	return assert
}

func (assert *boolAssertImpl) IsFalse() BoolAssert {
	return assert.isTrue(assert.actual == false,
		"Expected <false>, but was <%v>.", assert.actual)
}

func (assert *boolAssertImpl) IsTrue() BoolAssert {
	return assert.isTrue(assert.actual == true,
		"Expected <true>, but was <%v>.", assert.actual)
}
