package assert

type TestingT interface {
	Fail()
}

func Setup(t TestingT) AssertProvider {
	return setupImpl(t, theLogger)
}

type AssertProvider interface {
	That(actual interface{}) AnyTypeAssert
	ThatInt(actual int) IntAssert
	ThatString(actual string) StringAssert
	ThatBool(actual bool) BoolAssert
}
