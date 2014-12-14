package assert

func ThatString(t TestingT, actual string) *StringAssert {
	return &StringAssert{t, actual}
}

type StringAssert struct {
	t      TestingT
	actual string
}

func (assert *StringAssert) logf(format string, args ...interface{}) {
	assert.t.Errorf(format, args...)
}

func (assert *StringAssert) IsEqualTo(expected string) *StringAssert {
	if assert.actual != expected {
		assert.logf("Expected <%s>, but was <%s>.", expected, assert.actual)
	}
	return assert
}
