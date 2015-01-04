package assert

import (
	"reflect"
)

type AnyTypeAssert struct {
	logFacade *logFacade
	actual    interface{}
}

func (assert *AnyTypeAssert) isTrue(condition bool, format string, args ...interface{}) *AnyTypeAssert {
	logIfFalse(assert.logFacade, condition, format, args...)
	return assert
}

func (assert *AnyTypeAssert) IsEqualTo(expected interface{}) *AnyTypeAssert {
	return assert.isTrue(reflect.DeepEqual(assert.actual, expected),
		"Expected <%v>, but was <%v>.", expected, assert.actual)
}

func (assert *AnyTypeAssert) IsNotEqualTo(expected interface{}) *AnyTypeAssert {
	return assert.isTrue(!reflect.DeepEqual(assert.actual, expected),
		"Expected value not equal to <%v>, but was equal.", expected)
}

func (assert *AnyTypeAssert) IsNil() *AnyTypeAssert {
	return assert.isTrue(assert.actual == nil,
		"Expected value to be nil, but was <%v>.", assert.actual)
}

func (assert *AnyTypeAssert) IsNotNil() *AnyTypeAssert {
	return assert.isTrue(assert.actual != nil,
		"Expected value not to be nil, but was.")
}

func (assert *AnyTypeAssert) AsBool() *BoolAssert {
	val, kind := valueWithKind(assert.actual)
	if kind == reflect.Bool {
		return &BoolAssert{assert.logFacade, val.Bool()}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <bool>.", assert.actual, assert.actual)
	return &BoolAssert{}
}

func (assert *AnyTypeAssert) AsInt() *IntAssert {
	val, kind := valueWithKind(assert.actual)
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return &IntAssert{assert.logFacade, int(val.Int())}
	case reflect.Uint8, reflect.Uint16:
		return &IntAssert{assert.logFacade, int(val.Uint())}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <int>.", assert.actual, assert.actual)
	return &IntAssert{}
}

func (assert *AnyTypeAssert) AsInt64() *Int64Assert {
	val, kind := valueWithKind(assert.actual)
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &Int64Assert{assert.logFacade, val.Int()}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return &Int64Assert{assert.logFacade, int64(val.Uint())}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <int64>.", assert.actual, assert.actual)
	return &Int64Assert{}
}

func (assert *AnyTypeAssert) AsUint64() *Uint64Assert {
	val, kind := valueWithKind(assert.actual)
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &Uint64Assert{assert.logFacade, val.Uint()}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <uint64>.", assert.actual, assert.actual)
	return &Uint64Assert{}
}

func (assert *AnyTypeAssert) AsFloat() *FloatAssert {
	val, kind := valueWithKind(assert.actual)
	switch kind {
	case reflect.Float32, reflect.Float64:
		return &FloatAssert{assert.logFacade, val.Float()}
	case reflect.Int8, reflect.Int16, reflect.Int32:
		return &FloatAssert{assert.logFacade, float64(val.Int())}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return &FloatAssert{assert.logFacade, float64(val.Uint())}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <float64>.", assert.actual, assert.actual)
	return &FloatAssert{}
}

func (assert *AnyTypeAssert) AsComplex() *ComplexAssert {
	val, kind := valueWithKind(assert.actual)
	switch kind {
	case reflect.Complex64, reflect.Complex128:
		return &ComplexAssert{assert.logFacade, val.Complex()}
	case reflect.Float32, reflect.Float64:
		return &ComplexAssert{assert.logFacade, complex(val.Float(), 0)}
	case reflect.Int8, reflect.Int16, reflect.Int32:
		return &ComplexAssert{assert.logFacade, complex(float64(val.Int()), 0)}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return &ComplexAssert{assert.logFacade, complex(float64(val.Uint()), 0)}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <complex128>.", assert.actual, assert.actual)
	return &ComplexAssert{}
}

func (assert *AnyTypeAssert) AsString() *StringAssert {
	val, kind := valueWithKind(assert.actual)
	if kind == reflect.String {
		return &StringAssert{assert.logFacade, val.String()}
	}
	assert.isTrue(false, "Cannot convert <%v> of type <%T> to <string>.", assert.actual, assert.actual)
	return &StringAssert{}
}

func valueWithKind(data interface{}) (val reflect.Value, kind reflect.Kind) {
	if data != nil {
		val = reflect.ValueOf(data)
		kind = val.Type().Kind()
	}
	return
}
