package assert

func ThatString(t TestingT, actual string) *StringAssert {
	return &StringAssert{t, actual}
}

type StringAssert struct {
	t      TestingT
	actual string
}

func (assert *StringAssert) isTrue(condition bool, format string, args ...interface{}) *StringAssert {
	if !condition {
		assert.t.Errorf(format, args...)
	}
	return assert
}

func (assert *StringAssert) IsEqualTo(expected string) *StringAssert {
	return assert.isTrue(assert.actual == expected,
		"Expected <%s>, but was <%s>.", expected, assert.actual)
}
