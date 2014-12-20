package assert

func Setup(t TestingT) AssertProvider {
	return &assertProviderImpl{t, newLogger()}
}

type AssertProvider interface {
	ThatInt(actual int) IntAssert
	ThatString(actual string) StringAssert
	ThatBool(actual bool) BoolAssert
}
