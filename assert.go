package assert

func Setup(t TestingT) AssertProvider {
	return setupImpl(t, theLogger)
}

type AssertProvider interface {
	ThatInt(actual int) IntAssert
	ThatString(actual string) StringAssert
	ThatBool(actual bool) BoolAssert
}
