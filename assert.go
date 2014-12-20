package assert

func Setup(t TestingT) AssertProvider {
	return setupImpl(t, newLogger())
}

type AssertProvider interface {
	ThatInt(actual int) IntAssert
	ThatString(actual string) StringAssert
	ThatBool(actual bool) BoolAssert
}
