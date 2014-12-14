package assert

func ThatInt(t TestingT, actual int) *IntAssert {
	return &IntAssert{t, actual}
}

type IntAssert struct {
	t      TestingT
	actual int
}

func (assert *IntAssert) IsZero() *IntAssert {
	return assert.IsEqualTo(0)
}

func (assert *IntAssert) IsEqualTo(expected int) *IntAssert {
	if assert.actual != expected {
		assert.t.Errorf("Expected <%d>, but was <%d>.", expected, assert.actual)
	}
	return assert
}

func (assert *IntAssert) IsPositive() *IntAssert {
	if assert.actual <= 0 {
		assert.t.Errorf("Expected positive integer, but was <%d>.", assert.actual)
	}
	return assert
}

func (assert *IntAssert) IsNegative() *IntAssert {
	if assert.actual >= 0 {
		assert.t.Errorf("Expected negative integer, but was <%d>.", assert.actual)
	}
	return assert
}

func (assert *IntAssert) IsGreaterThan(smallerOrEqual int) *IntAssert {
	if assert.actual <= smallerOrEqual {
		assert.t.Errorf("Expected integer greater than <%d>, but was <%d>.", smallerOrEqual, assert.actual)
	}
	return assert
}

func (assert *IntAssert) IsGreaterOrEqualTo(smaller int) *IntAssert {
	if assert.actual < smaller {
		assert.t.Errorf("Expected integer greater or equal to <%d>, but was <%d>.", smaller, assert.actual)
	}
	return assert
}
