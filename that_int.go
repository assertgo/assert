package assert

func ThatInt(t TestingT, actual int) *IntAssert {
	return &IntAssert{t, actual}
}

type IntAssert struct {
	t      TestingT
	actual int
}

func (assert *IntAssert) logf(format string, args ...interface{}) {
	assert.t.Errorf(format, args...)
}

func (assert *IntAssert) IsZero() *IntAssert {
	return assert.IsEqualTo(0)
}

func (assert *IntAssert) IsEqualTo(expected int) *IntAssert {
	if assert.actual != expected {
		assert.logf("Expected <%d>, but was <%d>.", expected, assert.actual)
	}
	return assert
}

func (assert *IntAssert) IsPositive() *IntAssert {
	if assert.actual <= 0 {
		assert.logf("Expected positive integer, but was <%d>.", assert.actual)
	}
	return assert
}

func (assert *IntAssert) IsNegative() *IntAssert {
	if assert.actual >= 0 {
		assert.logf("Expected negative integer, but was <%d>.", assert.actual)
	}
	return assert
}

func (assert *IntAssert) IsGreaterThan(smallerOrEqual int) *IntAssert {
	if assert.actual <= smallerOrEqual {
		assert.logf("Expected integer greater than <%d>, but was <%d>.", smallerOrEqual, assert.actual)
	}
	return assert
}

func (assert *IntAssert) IsGreaterOrEqualTo(smaller int) *IntAssert {
	if assert.actual < smaller {
		assert.logf("Expected integer greater or equal to <%d>, but was <%d>.", smaller, assert.actual)
	}
	return assert
}
