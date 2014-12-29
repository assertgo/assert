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

func (assert *anyTypeAssertImpl) AsBool() BoolAssert {
	if actual, ok := assert.actual.(bool); ok {
		return &boolAssertImpl{assert.logFacade, actual}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <bool>.", assert.actual, assert.actual)
	return &boolAssertImpl{}
}

func (assert *anyTypeAssertImpl) AsInt() IntAssert {
	if assert.actual != nil {
		val := reflect.ValueOf(assert.actual)
		kind := val.Type().Kind()
		if kind == reflect.Int || kind == reflect.Int8 || kind == reflect.Int16 || kind == reflect.Int32 {
			return &intAssertImpl{assert.logFacade, int(val.Int())}
		} else if kind == reflect.Uint8 || kind == reflect.Uint16 {
			return &intAssertImpl{assert.logFacade, int(val.Uint())}
		}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <int>.", assert.actual, assert.actual)
	return &intAssertImpl{}
}

func (assert *anyTypeAssertImpl) AsString() StringAssert {
	if actual, ok := assert.actual.(string); ok {
		return &stringAssertImpl{assert.logFacade, actual}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <string>.", assert.actual, assert.actual)
	return &stringAssertImpl{}
}
