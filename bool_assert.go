package assert

func ThatBool(t TestingT, actual bool) BoolAssert {
	return &boolAssertImpl{t, actual}
}

type BoolAssert interface {
	IsFalse() BoolAssert
	IsTrue() BoolAssert
}
