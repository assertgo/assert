package assert

import "fmt"

type assertProviderImpl struct {
	logFacade *logFacade
}

type logFacade struct {
	t      TestingT
	logger errorLogger
}

func setupImpl(t TestingT, logger errorLogger) AssertProvider {
	return &assertProviderImpl{&logFacade{t, logger}}
}

func (assertProvider *assertProviderImpl) That(actual interface{}) AnyTypeAssert {
	return &anyTypeAssertImpl{assertProvider.logFacade, actual}
}

func (assertProvider *assertProviderImpl) ThatInt(actual int) IntAssert {
	return &intAssertImpl{assertProvider.logFacade, actual}
}

func (assertProvider *assertProviderImpl) ThatFloat(actual float64) FloatAssert {
	return &floatAssertImpl{assertProvider.logFacade, actual}
}

func (assertProvider *assertProviderImpl) ThatComplex(actual complex128) ComplexAssert {
	return &complexAssertImpl{assertProvider.logFacade, actual}
}

func (assertProvider *assertProviderImpl) ThatString(actual string) StringAssert {
	return &stringAssertImpl{assertProvider.logFacade, actual}
}

func (assertProvider *assertProviderImpl) ThatBool(actual bool) BoolAssert {
	return &boolAssertImpl{assertProvider.logFacade, actual}
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
