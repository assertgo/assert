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
	val, kind := valueWithKind(assert.actual)
	if kind == reflect.Bool {
		return &boolAssertImpl{assert.logFacade, val.Bool()}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <bool>.", assert.actual, assert.actual)
	return &boolAssertImpl{}
}

func (assert *anyTypeAssertImpl) AsInt() IntAssert {
	val, kind := valueWithKind(assert.actual)
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return &intAssertImpl{assert.logFacade, int(val.Int())}
	case reflect.Uint8, reflect.Uint16:
		return &intAssertImpl{assert.logFacade, int(val.Uint())}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <int>.", assert.actual, assert.actual)
	return &intAssertImpl{}
}

func (assert *anyTypeAssertImpl) AsString() StringAssert {
	val, kind := valueWithKind(assert.actual)
	if kind == reflect.String {
		return &stringAssertImpl{assert.logFacade, val.String()}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <string>.", assert.actual, assert.actual)
	return &stringAssertImpl{}
}

func valueWithKind(data interface{}) (val reflect.Value, kind reflect.Kind) {
	if data != nil {
		val = reflect.ValueOf(data)
		kind = val.Type().Kind()
	}
	return
}
