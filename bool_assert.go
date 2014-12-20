package assert

func ThatBool(t TestingT, actual bool) BoolAssert {
	return &boolAssertImpl{t, theLogger, actual}
}

type BoolAssert interface {
	IsFalse() BoolAssert
	IsTrue() BoolAssert
}
