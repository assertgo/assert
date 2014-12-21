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
