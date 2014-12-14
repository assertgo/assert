package assert

func ThatString(t TestingT, actual string) *StringAssertImpl {
	return &StringAssertImpl{t, actual}
}

type StringAssertImpl struct {
	t      TestingT
	actual string
}

func (assert *StringAssertImpl) isTrue(condition bool, format string, args ...interface{}) *StringAssertImpl {
	if !condition {
		assert.t.Errorf(format, args...)
	}
	return assert
}

func (assert *StringAssertImpl) IsEqualTo(expected string) *StringAssertImpl {
	return assert.isTrue(assert.actual == expected,
		"Expected <%s>, but was <%s>.", expected, assert.actual)
}
