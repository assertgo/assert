package assert

type T interface {
	Errorf(format string, args ...interface{})
}

func ThatInt(t T, actual int) *IntAssert {
	return &IntAssert{t, actual}
}

type IntAssert struct {
	t      T
	actual int
}

func (assert *IntAssert) IsZero() *IntAssert {
	return assert.IsEqualTo(0)
}

func (assert *IntAssert) IsEqualTo(expected int) *IntAssert {
	if assert.actual != expected {
		assert.t.Errorf("Expected <%d>, but was <%d>.\n", expected, assert.actual)
	}
	return assert
}
