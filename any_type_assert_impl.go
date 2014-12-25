package assert

import (
	"reflect"
)

type anyTypeAssertImpl struct {
	logFacade *logFacade
	actual    interface{}
}

func (assert *anyTypeAssertImpl) isTrue(condition bool, format string, args ...interface{}) *anyTypeAssertImpl {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *anyTypeAssertImpl) IsEqualTo(expected interface{}) AnyTypeAssert {
	return assert.isTrue(reflect.DeepEqual(assert.actual, expected),
		"Expected <%v>, but was <%v>.", expected, assert.actual)
}

func (assert *anyTypeAssertImpl) IsNotEqualTo(expected interface{}) AnyTypeAssert {
	return assert.isTrue(!reflect.DeepEqual(assert.actual, expected),
		"Expected value not equal to <%v>, but was equal.", expected)
}

func (assert *anyTypeAssertImpl) IsNil() AnyTypeAssert {
	return assert.isTrue(assert.actual == nil,
		"Expected value to be nil, but was <%v>.", assert.actual)
}

func (assert *anyTypeAssertImpl) IsNotNil() AnyTypeAssert {
	return assert.isTrue(assert.actual != nil,
		"Expected value not to be nil, but was.")
}

func (assert *anyTypeAssertImpl) AsInt() IntAssert {
	if actual, ok := assert.actual.(int); ok {
		return &intAssertImpl{assert.logFacade, actual}
	} else {
		assert.isTrue(false, "Cannot convert <%v> of type <%T> to <int>.", assert.actual, assert.actual)
		return &dummyIntAssertImpl{}
	}
}
