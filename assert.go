package assert

type TestingT interface {
	Fail()
}

func Setup(t TestingT) *Provider {
	return setupImpl(t, theLogger)
}
