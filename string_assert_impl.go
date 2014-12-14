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

func (assert *stringAssertImpl) IsNotEqualTo(nonexpected string) StringAssert {
	return assert.isTrue(assert.actual != nonexpected,
		"Expected string not equal to <%s>, but was equal.", nonexpected)
}

func (assert *stringAssertImpl) IsEmpty() StringAssert {
	return assert.isTrue(assert.actual == "",
		"Expected string to be empty, but was <%s>.", assert.actual)
}

func (assert *stringAssertImpl) IsInSlice(expectedSlice []string) StringAssert {
	return assert.isTrue(stringIsInSlice(expectedSlice, assert.actual),
		"Expected string to be in slice <%v>, but wasn't.", expectedSlice)
}

func stringIsInSlice(slice []string, expectedString string) bool {
	for _, v := range slice {
		if v == expectedString {
			return true
		}
	}
	return false
}
