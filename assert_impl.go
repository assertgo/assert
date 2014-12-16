package assert

type assertProviderImpl struct {
	t TestingT
}

func (assertProvider *assertProviderImpl) ThatInt(actual int) IntAssert {
	return ThatInt(assertProvider.t, actual)
}

func (assertProvider *assertProviderImpl) ThatString(actual string) StringAssert {
	return ThatString(assertProvider.t, actual)
}

func (assertProvider *assertProviderImpl) ThatBool(actual bool) BoolAssert {
	return ThatBool(assertProvider.t, actual)
}
