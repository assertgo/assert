package assert

type stringAssertImpl struct {
	t      TestingT
	actual string
}

func (assert *stringAssertImpl) isTrue(condition bool, format string, args ...interface{}) *stringAssertImpl {
	if !condition {
		assert.t.Errorf(format, args...)
	}
	return assert
}

func (assert *stringAssertImpl) IsEqualTo(expected string) StringAssert {
	return assert.isTrue(assert.actual == expected,
		"Expected <%s>, but was <%s>.", expected, assert.actual)
}
