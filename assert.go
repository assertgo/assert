package assert

type TestingT interface {
	Fail()
}

func Setup(t TestingT) *AssertProvider {
	return setupImpl(t, theLogger)
}
