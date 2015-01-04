package assert

import "fmt"

type AssertProvider struct {
	logFacade *logFacade
}

type logFacade struct {
	t      TestingT
	logger errorLogger
}

func setupImpl(t TestingT, logger errorLogger) *AssertProvider {
	return &AssertProvider{&logFacade{t, logger}}
}

func (assertProvider *AssertProvider) That(actual interface{}) *AnyTypeAssert {
	return &AnyTypeAssert{assertProvider.logFacade, actual}
}

func (assertProvider *AssertProvider) ThatInt(actual int) *IntAssert {
	return &IntAssert{assertProvider.logFacade, actual}
}

func (assertProvider *AssertProvider) ThatInt64(actual int64) *Int64Assert {
	return &Int64Assert{assertProvider.logFacade, actual}
}

func (assertProvider *AssertProvider) ThatUint64(actual uint64) *Uint64Assert {
	return &Uint64Assert{assertProvider.logFacade, actual}
}

func (assertProvider *AssertProvider) ThatFloat(actual float64) *FloatAssert {
	return &FloatAssert{assertProvider.logFacade, actual}
}

func (assertProvider *AssertProvider) ThatComplex(actual complex128) *ComplexAssert {
	return &ComplexAssert{assertProvider.logFacade, actual}
}

func (assertProvider *AssertProvider) ThatString(actual string) *StringAssert {
	return &StringAssert{assertProvider.logFacade, actual}
}

func (assertProvider *AssertProvider) ThatBool(actual bool) *BoolAssert {
	return &BoolAssert{assertProvider.logFacade, actual}
}

func logIfFalse(logFacade *logFacade, condition bool, format string, args ...interface{}) {
	if logFacade != nil && !condition {
		location := provideLocation(3)
		message := fmt.Sprintf(format, args...)
		logFacade.Log(location, message)
	}
}

func (logFacade *logFacade) Log(location *location, message string) {
	logFacade.t.Fail()
	logFacade.logger.Log(location, message)
}
