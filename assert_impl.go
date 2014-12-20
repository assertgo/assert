package assert

type assertProviderImpl struct {
	t TestingT
}

func (assertProvider *assertProviderImpl) ThatInt(actual int) IntAssert {
	return &intAssertImpl{assertProvider.t, actual}
}

func (assertProvider *assertProviderImpl) ThatString(actual string) StringAssert {
	return &stringAssertImpl{assertProvider.t, actual}
}

func (assertProvider *assertProviderImpl) ThatBool(actual bool) BoolAssert {
	return &boolAssertImpl{assertProvider.t, actual}
}
