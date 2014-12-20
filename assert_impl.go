package assert

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

func (logFacade *logFacade) Log(location *location, message string) {
	logFacade.t.Fail()
	logFacade.logger.Log(location, message)
}
