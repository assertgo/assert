package assert

import (
	"reflect"
)

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
