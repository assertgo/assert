package assert

type assertProviderImpl struct {
	t      TestingT
	logger errorLogger
}

func setupImpl(t TestingT, logger errorLogger) AssertProvider {
	return &assertProviderImpl{t, logger}
}

func (assertProvider *assertProviderImpl) ThatInt(actual int) IntAssert {
	return &intAssertImpl{assertProvider.t, assertProvider.logger, actual}
}

func (assertProvider *assertProviderImpl) ThatString(actual string) StringAssert {
	return &stringAssertImpl{assertProvider.t, assertProvider.logger, actual}
}

func (assertProvider *assertProviderImpl) ThatBool(actual bool) BoolAssert {
	return &boolAssertImpl{assertProvider.t, assertProvider.logger, actual}
}
