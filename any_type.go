package assert

import "reflect"

type AnyType struct {
	logFacade *logFacade
	actual    interface{}
}

func (a *AnyType) IsEqualTo(expected interface{}) *AnyType {
	return a.isTrue(reflect.DeepEqual(a.actual, expected),
		"Expected <%v>, but was <%v>.", expected, a.actual)
}

func (a *AnyType) IsNotEqualTo(expected interface{}) *AnyType {
	return a.isTrue(!reflect.DeepEqual(a.actual, expected),
		"Expected value not equal to <%v>, but was equal.", expected)
}

func (a *AnyType) IsNil() *AnyType {
	return a.isTrue(a.actual == nil,
		"Expected value to be nil, but was <%v>.", a.actual)
}

func (a *AnyType) IsNotNil() *AnyType {
	return a.isTrue(a.actual != nil,
		"Expected value not to be nil, but was.")
}

func (a *AnyType) AsBool() *Bool {
	val, kind := valueWithKind(a.actual)
	if kind == reflect.Bool {
		return &Bool{a.logFacade, val.Bool()}
	}
	a.isTrue(false, "Cannot convert <%v> of type <%T> to <bool>.", a.actual, a.actual)
	return &Bool{}
}

func (a *AnyType) AsInt() *Int {
	val, kind := valueWithKind(a.actual)
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		return &Int{a.logFacade, int(val.Int())}
	case reflect.Uint8, reflect.Uint16:
		return &Int{a.logFacade, int(val.Uint())}
	}
	a.isTrue(false, "Cannot convert <%v> of type <%T> to <int>.", a.actual, a.actual)
	return &Int{}
}

func (a *AnyType) AsInt64() *Int64 {
	val, kind := valueWithKind(a.actual)
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &Int64{a.logFacade, val.Int()}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return &Int64{a.logFacade, int64(val.Uint())}
	}
	a.isTrue(false, "Cannot convert <%v> of type <%T> to <int64>.", a.actual, a.actual)
	return &Int64{}
}

func (a *AnyType) AsUint64() *Uint64 {
	val, kind := valueWithKind(a.actual)
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &Uint64{a.logFacade, val.Uint()}
	}
	a.isTrue(false, "Cannot convert <%v> of type <%T> to <uint64>.", a.actual, a.actual)
	return &Uint64{}
}

func (a *AnyType) AsFloat() *Float {
	val, kind := valueWithKind(a.actual)
	switch kind {
	case reflect.Float32, reflect.Float64:
		return &Float{a.logFacade, val.Float()}
	case reflect.Int8, reflect.Int16, reflect.Int32:
		return &Float{a.logFacade, float64(val.Int())}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return &Float{a.logFacade, float64(val.Uint())}
	}
	a.isTrue(false, "Cannot convert <%v> of type <%T> to <float64>.", a.actual, a.actual)
	return &Float{}
}

func (a *AnyType) AsComplex() *Complex {
	val, kind := valueWithKind(a.actual)
	switch kind {
	case reflect.Complex64, reflect.Complex128:
		return &Complex{a.logFacade, val.Complex()}
	case reflect.Float32, reflect.Float64:
		return &Complex{a.logFacade, complex(val.Float(), 0)}
	case reflect.Int8, reflect.Int16, reflect.Int32:
		return &Complex{a.logFacade, complex(float64(val.Int()), 0)}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return &Complex{a.logFacade, complex(float64(val.Uint()), 0)}
	}
	a.isTrue(false, "Cannot convert <%v> of type <%T> to <complex128>.", a.actual, a.actual)
	return &Complex{}
}

func (a *AnyType) AsString() *String {
	val, kind := valueWithKind(a.actual)
	if kind == reflect.String {
		return &String{a.logFacade, val.String()}
	}
	a.isTrue(false, "Cannot convert <%v> of type <%T> to <string>.", a.actual, a.actual)
	return &String{}
}

func (a *AnyType) isTrue(condition bool, format string, args ...interface{}) *AnyType {
	logIfFalse(a.logFacade, condition, format, args...)
	return a
}

func valueWithKind(data interface{}) (val reflect.Value, kind reflect.Kind) {
	if data != nil {
		val = reflect.ValueOf(data)
		kind = val.Type().Kind()
	}
	return
}
