package assert

func Setup(t TestingT) AssertProvider {
	return &assertProviderImpl{t}
}

type AssertProvider interface {
	ThatInt(actual int) IntAssert
	ThatString(actual string) StringAssert
}

type assertProviderImpl struct {
	t TestingT
}

func (assertProvider *assertProviderImpl) ThatInt(actual int) IntAssert {
	return ThatInt(assertProvider.t, actual)
}

func (assertProvider *assertProviderImpl) ThatString(actual string) StringAssert {
	return ThatString(assertProvider.t, actual)
}
