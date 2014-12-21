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

func (assertProvider *assertProviderImpl) ThatInt(actual int) IntAssert {
	return &intAssertImpl{assertProvider.logFacade, actual}
}

func (assertProvider *assertProviderImpl) ThatString(actual string) StringAssert {
	return &stringAssertImpl{assertProvider.logFacade, actual}
}

func (assertProvider *assertProviderImpl) ThatBool(actual bool) BoolAssert {
	return &boolAssertImpl{assertProvider.logFacade, actual}
}

func logIfFalse(logFacade *logFacade, condition bool, format string, args ...interface{}) {
	if !condition {
		location := provideLocation(3)
		message := fmt.Sprintf(format, args...)
		logFacade.Log(&location, message)
	}
}

func (logFacade *logFacade) Log(location *location, message string) {
	logFacade.t.Fail()
	logFacade.logger.Log(location, message)
}
