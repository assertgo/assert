package assert

func Setup(t TestingT) AssertProvider {
	return &assertProviderImpl{t}
}

type AssertProvider interface {
	ThatInt(actual int) IntAssert
	ThatString(actual string) StringAssert
	ThatBool(actual bool) BoolAssert
}
